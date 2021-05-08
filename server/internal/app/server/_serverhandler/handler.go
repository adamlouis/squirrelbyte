package serverhandler

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/internal/app/server/serverdef"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/document"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/document/documentsqlite3"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/job"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/job/jobsqlite3"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/oauth"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/secret"
	"github.com/jmoiron/sqlx"
)

// break apart job & documents
// NewAPIHandler returns an implementation of the APIHandler interface
func NewAPIHandler(
	db *sqlx.DB,
) (serverdef.APIHandler, error) {
	err := documentsqlite3.NewDocumentRepository(db).Init(context.Background())
	if err != nil {
		return nil, err
	}

	err = jobsqlite3.NewJobRepository(db).Init(context.Background())
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
	Document       document.Repository
	Job            job.Repository
	SecretsManager secret.Manager
	OauthManager   oauth.Manager
	Commit         func() error
	Rollback       func() error
}

func (a *apiHandler) GetRepositories() (*Repositories, error) {
	tx, err := a.db.Beginx()
	if err != nil {
		return nil, err
	}

	dr := documentsqlite3.NewDocumentRepository(tx)
	jr := jobsqlite3.NewJobRepository(tx)

	return &Repositories{
		Document:       dr,
		Job:            jr,
		SecretsManager: secret.NewManger(),
		OauthManager:   oauth.NewManger(),
		Commit:         tx.Commit,
		Rollback:       tx.Rollback,
	}, nil
}
