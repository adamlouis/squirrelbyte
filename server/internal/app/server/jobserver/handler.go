package jobserver

import (
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/job"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/job/jobsqlite3"
	"github.com/jmoiron/sqlx"
)

func NewAPIHandler(db *sqlx.DB) APIHandler {
	return &hdl{
		db: db,
	}
}

type hdl struct {
	db *sqlx.DB
}

type CommitFn func() error
type RollbackFn func() error

func (h *hdl) GetRepository() (job.Repository, CommitFn, RollbackFn, error) {
	tx, err := h.db.Beginx()
	if err != nil {
		return nil, nil, nil, err
	}
	return jobsqlite3.NewJobRepository(tx), tx.Commit, tx.Rollback, nil
}
