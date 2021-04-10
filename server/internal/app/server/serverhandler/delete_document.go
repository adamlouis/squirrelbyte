package serverhandler

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/internal/app/server/serverdef"
)

func (a *apiHandler) DeleteDocument(ctx context.Context, pathParams *serverdef.DeleteDocumentPathParams) error {
	repos, err := a.GetRepositories()
	if err != nil {
		return err
	}
	defer repos.Rollback() //nolint

	err = repos.Document.Delete(ctx, pathParams.DocumentID)
	if err != nil {
		return err
	}

	return repos.Commit()
}
