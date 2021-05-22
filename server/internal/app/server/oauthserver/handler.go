package oauthserver

import (
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/oauth"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/oauth/oauthsqlite3"
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

func (h *hdl) GetRepository() (oauth.Repository, CommitFn, RollbackFn, error) {
	tx, err := h.db.Beginx()
	if err != nil {
		return nil, nil, nil, err
	}
	return oauthsqlite3.NewRepository(tx), tx.Commit, tx.Rollback, nil
}
