package documentserver

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/document"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/document/documentsqlite3"
	"github.com/jmoiron/sqlx"
)

func NewAPIHandler(db *sqlx.DB) APIHandler {
	documentsqlite3.NewDocumentRepository(db).Init(context.Background())

	return &hdl{
		db: db,
	}
}

type hdl struct {
	db *sqlx.DB
}

type CommitFn func() error
type RollbackFn func() error

func (h *hdl) GetRepository() (document.Repository, CommitFn, RollbackFn, error) {
	tx, err := h.db.Beginx()
	if err != nil {
		return nil, nil, nil, err
	}
	return documentsqlite3.NewDocumentRepository(tx), tx.Commit, tx.Rollback, nil
}
