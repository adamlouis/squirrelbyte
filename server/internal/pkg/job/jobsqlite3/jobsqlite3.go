package jobsqlite3

import (
	"context"
	"errors"
	"fmt"
	"time"

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
		output TEXT NULL CHECK(json_type(output) == 'object'),
		succeed_at TEXT NULL,
		errored_at TEXT NULL,
		created_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL CHECK(id <> ''),
		updated_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL CHECK(id <> '')
	);

	CREATE INDEX IF NOT EXISTS job_name ON job(name);
	CREATE INDEX IF NOT EXISTS job_status ON job(status);
	CREATE INDEX IF NOT EXISTS job_created_at ON job(created_at);

	CREATE TRIGGER IF NOT EXISTS set_job_updated_at
	AFTER UPDATE ON job
	BEGIN
		UPDATE job SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
	END;`)
	return err
}

type jobRow struct {
	ID          string  `db:"id"`
	Name        string  `db:"name"`
	Status      string  `db:"status"`
	Input       []byte  `db:"input"`
	Output      *[]byte `db:"output"`
	SucceededAt *string `db:"succeed_at"`
	ErroredAt   *string `db:"errored_at"`
	CreatedAt   string  `db:"created_at"`
	UpdatedAt   string  `db:"updated_at"`
}

func (jr *jobRepo) Queue(ctx context.Context, j *job.Job) (*job.Job, error) {
	j.ID = uuid.New().String()
	j.Status = job.JobStatusQueued

	_, err := jr.db.Exec(`
			INSERT INTO
				job
					(id, name, status, input)
				VALUES
					(?, ?, ?, ?)`,
		j.ID, j.Name, j.Status, j.Input)
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
			id, name, status, input, output, succeed_at, errored_at, created_at, updated_at
		FROM job
		WHERE id = ?`,
		id,
	)

	var r jobRow
	err := row.StructScan(&r)
	if err != nil {
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
		id, name, status, input, output, succeed_at, errored_at, created_at, updated_at
	FROM job
	ORDER BY created_at`,
	)

	if err != nil {
		return nil, err
	}

	jobs := []*job.Job{}
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
	row := jr.db.QueryRowx(`
		SELECT id, status
		FROM job
		WHERE status = ?
		ORDER BY created_at
		LIMIT 1`,
		job.JobStatusQueued,
	)

	var r jobRow
	err := row.StructScan(&r)
	if err != nil {
		return nil, err
	}

	r.Status = job.JobStatusClaimed
	_, err = jr.db.Exec(`UPDATE job SET status = ? WHERE id = ?`, r.Status, r.ID)
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
	_, err = jr.db.Exec(`UPDATE job SET status = ? WHERE id = ?`, job.JobStatusQueued, j.ID)
	if err != nil {
		return nil, err
	}

	return jr.Get(ctx, id)
}
func (jr *jobRepo) Success(ctx context.Context, id string, out interface{}) (*job.Job, error) {
	j, err := jr.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if j.Status != job.JobStatusClaimed {
		return nil, fmt.Errorf("only jobs with status CLAIMED can be updated with success - %s has status %s", j.ID, j.Status)
	}

	j.Status = job.JobStatusSuccess
	_, err = jr.db.Exec(`UPDATE job SET status = ?, output = ?, succeed_at = CURRENT_TIMESTAMP WHERE id = ?`, job.JobStatusSuccess, out, j.ID)
	if err != nil {
		return nil, err
	}

	return jr.Get(ctx, id)
}
func (jr *jobRepo) Error(ctx context.Context, id string, out interface{}) (*job.Job, error) {
	j, err := jr.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if j.Status != job.JobStatusClaimed {
		return nil, fmt.Errorf("only jobs with status CLAIMED can be update with error - %s has status %s", j.ID, j.Status)
	}

	j.Status = job.JobStatusError
	_, err = jr.db.Exec(`UPDATE job SET status = ?, output = ?, errored_at = CURRENT_TIMESTAMP WHERE id = ?`, job.JobStatusError, out, j.ID)
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

	var s *time.Time
	if r.SucceededAt != nil {
		t, err := time.Parse(datetimeFormat, *r.SucceededAt)
		if err != nil {
			return nil, err
		}
		s = &t
	}

	var e *time.Time
	if r.ErroredAt != nil {
		t, err := time.Parse(datetimeFormat, *r.ErroredAt)
		if err != nil {
			return nil, err
		}
		e = &t
	}

	return &job.Job{
		ID:          r.ID,
		Name:        r.Name,
		Status:      job.JobStatus(r.Status),
		Input:       r.Input,
		Output:      r.Output,
		SucceededAt: s,
		ErroredAt:   e,
		CreatedAt:   c,
		UpdatedAt:   u,
	}, nil
}