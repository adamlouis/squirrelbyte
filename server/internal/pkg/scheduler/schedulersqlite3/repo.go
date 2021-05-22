package schedulersqlite3

import (
	"context"
	"database/sql"
	"embed"
	"encoding/json"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/crudutil"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/errtype"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/scheduler"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/sqlite3util"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/schedulermodel"
	"github.com/jmoiron/sqlx"
)

func NewRepo(db sqlx.Ext) scheduler.Repository {
	return &repo{db: db}
}

//go:embed migration/*.sql
var MigrationFS embed.FS

type repo struct {
	db sqlx.Ext
}

type schedulerRow struct {
	ID        string `db:"id"`
	Schedule  string `db:"schedule"`
	JobName   string `db:"job_name"`
	Input     []byte `db:"input"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func (r *repo) Put(ctx context.Context, in *schedulermodel.Scheduler) (*schedulermodel.Scheduler, error) {
	input, err := json.Marshal(in.Input)
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(`
			INSERT INTO
				scheduler(
					id,
					schedule,
					job_name,
					input
				)
			VALUES
				(
					?,
					?,
					?,
					?
				)
			ON CONFLICT (id)
			DO UPDATE
			SET
				schedule = ?,
				job_name = ?,
				input = ?
		`,
		in.ID,
		in.Schedule, in.JobName, input,
		in.Schedule, in.JobName, input,
	)
	if err != nil {
		return nil, err
	}
	return r.Get(ctx, in.ID)
}
func (r *repo) Get(ctx context.Context, id string) (*schedulermodel.Scheduler, error) {
	row := r.db.QueryRowx(`
		SELECT
			id,
			schedule,
			job_name,
			input,
			created_at,
			updated_at
		FROM scheduler
		WHERE id = ?`,
		id,
	)

	var sr schedulerRow
	err := row.StructScan(&sr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errtype.NotFoundError{Err: err}
		}
		return nil, err
	}

	out, err := schedulerRowToScheduler(&sr)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *repo) Delete(ctx context.Context, id string) error {
	return crudutil.Delete(r.db, `DELETE FROM scheduler WHERE id = ?`, id)
}

func (r *repo) List(ctx context.Context, args *schedulermodel.ListSchedulersRequest) (*schedulermodel.ListSchedulersResponse, error) {
	sz, err := crudutil.GetPageSize(args.PageSize, 500)
	if err != nil {
		return nil, err
	}

	sb := sq.
		StatementBuilder.
		Select("id, schedule, job_name, input, created_at, updated_at").
		From("scheduler").
		OrderBy("id ASC").
		Limit(uint64(sz) + 1) // get n+1 so we know if there's a next page

	if args.PageToken != "" {
		page := &listSchedulersPageData{}
		err := crudutil.DecodePageData(args.PageToken, page)
		if err != nil {
			return nil, err
		}
		sb = sb.Where(sq.GtOrEq{"id": page.NextID})
	}

	sql, sqlArgs, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Queryx(sql, sqlArgs...)
	if err != nil {
		return nil, err
	}

	schedulers := make([]*schedulermodel.Scheduler, 0, sz)
	for rows.Next() {
		var r schedulerRow
		err = rows.StructScan(&r)
		if err != nil {
			return nil, err
		}
		s, err := schedulerRowToScheduler(&r)
		if err != nil {
			return nil, err
		}
		schedulers = append(schedulers, s)
	}

	nextPageToken := ""
	if len(schedulers) > int(sz) {
		s, err := crudutil.EncodePageData(&listSchedulersPageData{
			NextID: schedulers[len(schedulers)-1].ID,
		})
		if err != nil {
			return nil, err
		}
		nextPageToken = s
		schedulers = schedulers[0 : len(schedulers)-1]
	}

	return &schedulermodel.ListSchedulersResponse{
		Schedulers:    schedulers,
		NextPageToken: nextPageToken,
	}, nil
}

type listSchedulersPageData struct {
	NextID string `json:"next_id"`
}

func schedulerRowToScheduler(r *schedulerRow) (*schedulermodel.Scheduler, error) {
	c, err := time.Parse(sqlite3util.DatetimeFormat, r.CreatedAt)
	if err != nil {
		return nil, err
	}

	u, err := time.Parse(sqlite3util.DatetimeFormat, r.UpdatedAt)
	if err != nil {
		return nil, err
	}

	var input map[string]interface{}
	err = json.Unmarshal(r.Input, &input)
	if err != nil {
		return nil, err
	}
	return &schedulermodel.Scheduler{
		ID:        r.ID,
		Schedule:  r.Schedule,
		JobName:   r.JobName,
		Input:     input,
		CreatedAt: present.ToAPITime(c),
		UpdatedAt: present.ToAPITime(u),
	}, nil
}
