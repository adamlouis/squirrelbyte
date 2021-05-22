package kvserver

import (
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/kv"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/kv/kvsqlite3"
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

func (h *hdl) GetRepository() (kv.Repository, CommitFn, RollbackFn, error) {
	tx, err := h.db.Beginx()
	if err != nil {
		return nil, nil, nil, err
	}
	return kvsqlite3.NewRepository(tx), tx.Commit, tx.Rollback, nil
}
