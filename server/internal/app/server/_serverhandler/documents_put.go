package serverhandler

import (
	"context"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/internal/app/server/serverdef"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

func (a *apiHandler) PutDocument(ctx context.Context, pathParams *model.PutDocumentPathParams, d *model.Document) (*model.Document, error) {
	if pathParams.DocumentID != d.ID {
		return nil, serverdef.NewHTTPErrorFromString(http.StatusBadRequest, "document id in path does not match document id in request body")
	}
	repos, err := a.GetRepositories()
	if err != nil {
		return nil, err
	}
	defer repos.Rollback() //nolint

	e, err := present.APIDocumentToInternalDocument(d)
	if err != nil {
		return nil, err
	}

	created, err := repos.Document.Put(ctx, e)
	if err != nil {
		return nil, err
	}

	if err := repos.Commit(); err != nil {
		return nil, err
	}

	return present.InternalDocumentToAPIDocument(created)
}
