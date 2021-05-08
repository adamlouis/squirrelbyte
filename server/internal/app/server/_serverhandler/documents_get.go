package serverhandler

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

func (a *apiHandler) GetDocument(ctx context.Context, pathParams *model.GetDocumentPathParams) (*model.Document, error) {
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
