package jobsqlite3

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/job"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewJobRepository(db sqlx.Ext) job.Repository {
	return &jobRepo{
		db: db,
	}
}

type jobRepo struct {
	db sqlx.Ext
}

const (
	datetimeFormat = "2006-01-02 15:04:05" // note: no T
)

// TODO: rwmutex

func (jr *jobRepo) Init(ctx context.Context) error {
	// todo: real migration, not this string
	// migrate on startup
	// fine for now
	_, err := jr.db.Exec(`
	CREATE TABLE IF NOT EXISTS job(
		id TEXT NOT NULL UNIQUE CHECK(id <> ''),
		name TEXT NOT NULL CHECK(name <> ''),
		status TEXT NOT NULL CHECK(status <> ''),
		input TEXT NOT NULL CHECK(json_type(input) == 'object'),
		succeed_at TEXT NULL,
		errored_at TEXT NULL,
		claimed_at TEXT NULL,
		scheduled_for TEXT NULL,
		created_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL CHECK(id <> ''),
		updated_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL CHECK(id <> '')
	);

	CREATE INDEX IF NOT EXISTS job_name ON job(name);
	CREATE INDEX IF NOT EXISTS job_status ON job(status);
	CREATE INDEX IF NOT EXISTS job_created_at ON job(created_at);
	CREATE INDEX IF NOT EXISTS job_scheduled_for ON job(scheduled_for);

	CREATE TRIGGER IF NOT EXISTS set_job_updated_at
	AFTER UPDATE ON job
	BEGIN
		UPDATE job SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
	END;`)
	return err
}

type jobRow struct {
	ID           string  `db:"id"`
	Name         string  `db:"name"`
	Status       string  `db:"status"`
	Input        []byte  `db:"input"`
	SucceededAt  *string `db:"succeed_at"`
	ClaimedAt    *string `db:"claimed_at"`
	ScheduledFor *string `db:"scheduled_for"`
	ErroredAt    *string `db:"errored_at"`
	CreatedAt    string  `db:"created_at"`
	UpdatedAt    string  `db:"updated_at"`
}

func (jr *jobRepo) Queue(ctx context.Context, j *job.Job) (*job.Job, error) {
	j.ID = uuid.New().String()
	j.Status = job.JobStatusQueued

	var scheduledForStr *string
	if j.ScheduledFor != nil {
		s := j.ScheduledFor.Format(datetimeFormat)
		scheduledForStr = &s
	}

	_, err := jr.db.Exec(`
			INSERT INTO
				job
					(id, name, status, input, scheduled_for)
				VALUES
					(?, ?, ?, ?, ?)`,
		j.ID, j.Name, j.Status, j.Input, scheduledForStr)
	if err != nil {
		return nil, err
	}

	return jr.Get(ctx, j.ID)
}

func (jr *jobRepo) Delete(ctx context.Context, id string) error {
	r, err := jr.db.Exec(`DELETE FROM job WHERE id = ?`, id)

	if err != nil {
		return err
	}

	ct, err := r.RowsAffected()
	if err != nil {
		return err
	}

	if ct == 0 {
		return errors.New("not found")
	}
	if ct > 1 {
		return errors.New("unexpected")
	}

	return nil
}

func (jr *jobRepo) Get(ctx context.Context, id string) (*job.Job, error) {
	row := jr.db.QueryRowx(`
		SELECT
			id, name, status, input, succeed_at, errored_at, claimed_at, created_at, updated_at, scheduled_for
		FROM job
		WHERE id = ?`,
		id,
	)

	var r jobRow
	err := row.StructScan(&r)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("job %s not found", id) // TODO: 404
		}
		return nil, err
	}

	j, err := jobRowToJob(&r)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (jr *jobRepo) List(ctx context.Context, args *job.ListJobArgs) (*job.ListJobResults, error) {
	rows, err := jr.db.Queryx(`
	SELECT
		id, name, status, input, succeed_at, errored_at, claimed_at, created_at, updated_at, scheduled_for
	FROM job
	ORDER BY created_at`,
	)
	jobs := []*job.Job{}

	if err != nil {
		if err == sql.ErrNoRows {
			return &job.ListJobResults{Jobs: jobs}, nil
		}
		return nil, err
	}

	for rows.Next() {
		var r jobRow
		err = rows.StructScan(&r)
		if err != nil {
			return nil, err
		}
		j, err := jobRowToJob(&r)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, j)
	}
	return &job.ListJobResults{Jobs: jobs}, nil
}

// TODO - handle concurrency, locking, etc
func (jr *jobRepo) Claim(ctx context.Context, opts job.ClaimOptions) (*job.Job, error) {
	sb := sq.
		StatementBuilder.
		Select("id, status").
		From("job").
		OrderBy("created_at ASC").
		Where(sq.Eq{"status": job.JobStatusQueued}).
		Where(sq.Or{
			sq.Expr("scheduled_for IS NULL"),
			sq.LtOrEq{"scheduled_for": "CURRENT_TIMESTAMP"},
		}).
		Limit(1) // get n+1 so we know if there's a next page

	if opts.JobID != "" {
		sb = sb.Where(sq.Eq{"id": opts.JobID})
	}

	if len(opts.Names) > 0 {
		ors := make(sq.Or, len(opts.Names))
		for i, n := range opts.Names {
			ors[i] = sq.Eq{"name": n}
		}
		sb = sb.Where(ors)
	}

	query, queryArgs, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	row := jr.db.QueryRowx(query, queryArgs...)

	var r jobRow
	err = row.StructScan(&r)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	r.Status = string(job.JobStatusClaimed)
	_, err = jr.db.Exec(`UPDATE job SET status = ?, claimed_at = CURRENT_TIMESTAMP WHERE id = ?`, r.Status, r.ID)
	if err != nil {
		return nil, err
	}

	return jr.Get(ctx, r.ID)
}

func (jr *jobRepo) Release(ctx context.Context, id string) (*job.Job, error) {
	j, err := jr.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if j.Status != job.JobStatusClaimed {
		return nil, fmt.Errorf("only jobs with status CLAIMED can be released - %s has status %s", j.ID, j.Status)
	}

	j.Status = job.JobStatusQueued
	_, err = jr.db.Exec(`UPDATE job SET status = ?, claimed_at = NULL WHERE id = ?`, job.JobStatusQueued, j.ID)
	if err != nil {
		return nil, err
	}

	return jr.Get(ctx, id)
}
func (jr *jobRepo) Success(ctx context.Context, id string) (*job.Job, error) {
	j, err := jr.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if j.Status != job.JobStatusClaimed {
		return nil, fmt.Errorf("only jobs with status CLAIMED can be updated with success - %s has status %s", j.ID, j.Status)
	}

	j.Status = job.JobStatusSuccess
	_, err = jr.db.Exec(`UPDATE job SET status = ?, succeed_at = CURRENT_TIMESTAMP WHERE id = ?`, job.JobStatusSuccess, j.ID)
	if err != nil {
		return nil, err
	}

	return jr.Get(ctx, id)
}
func (jr *jobRepo) Error(ctx context.Context, id string) (*job.Job, error) {
	j, err := jr.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if j.Status != job.JobStatusClaimed {
		return nil, fmt.Errorf("only jobs with status CLAIMED can be update with error - %s has status %s", j.ID, j.Status)
	}

	j.Status = job.JobStatusError
	_, err = jr.db.Exec(`UPDATE job SET status = ?, errored_at = CURRENT_TIMESTAMP WHERE id = ?`, job.JobStatusError, j.ID)
	if err != nil {
		return nil, err
	}

	return jr.Get(ctx, id)
}

func jobRowToJob(r *jobRow) (*job.Job, error) {
	c, err := time.Parse(datetimeFormat, r.CreatedAt)
	if err != nil {
		return nil, err
	}

	u, err := time.Parse(datetimeFormat, r.UpdatedAt)
	if err != nil {
		return nil, err
	}

	sa, err := tptr(r.SucceededAt)
	if err != nil {
		return nil, err
	}

	er, err := tptr(r.ErroredAt)
	if err != nil {
		return nil, err
	}

	ca, err := tptr(r.ClaimedAt)
	if err != nil {
		return nil, err
	}

	sf, err := tptr(r.ScheduledFor)
	if err != nil {
		return nil, err
	}

	return &job.Job{
		ID:           r.ID,
		Name:         r.Name,
		Status:       job.JobStatus(r.Status),
		Input:        r.Input,
		SucceededAt:  sa,
		ErroredAt:    er,
		ClaimedAt:    ca,
		ScheduledFor: sf,
		CreatedAt:    c,
		UpdatedAt:    u,
	}, nil
}

func tptr(s *string) (*time.Time, error) {
	if s != nil {
		t, err := time.Parse(datetimeFormat, *s)
		if err != nil {
			return nil, err
		}
		return &t, nil
	}
	return nil, nil
}
