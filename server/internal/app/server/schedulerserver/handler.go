package schedulerserver

import (
	"context"
	"time"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/jsonlog"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/scheduler"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/scheduler/schedulersqlite3"
	"github.com/adamlouis/squirrelbyte/server/pkg/client/jobclient"
	"github.com/jmoiron/sqlx"
)

func NewAPIHandler(ctx context.Context, db *sqlx.DB, jobClient jobclient.Client) APIHandler {
	updatedIDChan := make(chan string)
	runnerErrChan := make(chan error)

	runner := scheduler.NewRunner(
		schedulersqlite3.NewRepo(db),
		jobClient,
		runnerErrChan,
	)

	go func() {
		for err := range runnerErrChan {
			jsonlog.Log(
				"name", "RunnerError",
				"error", err.Error(),
				"timestamp", time.Now(),
			)
		}
	}()

	go func() {
		for updatedID := range updatedIDChan {
			jsonlog.Log(
				"name", "RunnerUpdate",
				"scheduler_id", updatedID,
				"timestamp", time.Now(),
			)
			runner.Update(ctx, updatedID)
		}
	}()

	go func() {
		for {
			err := runner.Run(ctx)
			runnerErrChan <- err
			time.Sleep(5 * time.Second)
		}
	}()

	return &hdl{
		db:            db,
		updatedIDChan: updatedIDChan,
	}
}

type hdl struct {
	db            *sqlx.DB
	updatedIDChan chan string
}

type CommitFn func() error
type RollbackFn func() error

func (h *hdl) GetRepository() (scheduler.Repository, CommitFn, RollbackFn, error) {
	tx, err := h.db.Beginx()
	if err != nil {
		return nil, nil, nil, err
	}
	return schedulersqlite3.NewRepo(tx), tx.Commit, tx.Rollback, nil
}

func (h *hdl) onUpdate(id string) {
	h.updatedIDChan <- id
}
