package serverhandler

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/internal/app/server/serverdef"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
)

func (a *apiHandler) GetDocument(ctx context.Context, pathParams *serverdef.GetDocumentPathParams) (*serverdef.Document, error) {
	repos, err := a.GetRepositories()
	if err != nil {
		return nil, err
	}
	defer repos.Rollback() //nolint

	d, err := repos.Document.Get(ctx, pathParams.DocumentID)
	if err != nil {
		return nil, err
	}

	return present.InternalDocumentToAPIDocument(d)
}
