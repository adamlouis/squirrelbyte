package serverhandler

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/internal/app/server/serverdef"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/document"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/document/documentsqlite3"
	"github.com/jmoiron/sqlx"
)

// NewAPIHandler returns an implementation of the APIHandler interface
func NewAPIHandler(
	db *sqlx.DB,
) (serverdef.APIHandler, error) {
	err := documentsqlite3.NewDocumentRepository(db).Init(context.Background())

	if err != nil {
		return nil, err
	}

	return &apiHandler{
		db: db,
	}, nil
}

type apiHandler struct {
	db *sqlx.DB
}

// Repositories wraps the resource repositories for the server
type Repositories struct {
	Document document.Repository
	Commit   func() error
	Rollback func() error
}

func (a *apiHandler) GetRepositories() (*Repositories, error) {
	tx, err := a.db.Beginx()
	if err != nil {
		return nil, err
	}

	dr := documentsqlite3.NewDocumentRepository(tx)

	return &Repositories{
		Document: dr,
		Commit:   tx.Commit,
		Rollback: tx.Rollback,
	}, nil
}
