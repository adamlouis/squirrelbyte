package serverhandler

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/internal/app/server/serverdef"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/document"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
)

func (a *apiHandler) ListDocuments(ctx context.Context, queryParams *serverdef.ListDocumentsQueryParams) (*serverdef.ListDocumentsResponse, error) {
	repos, err := a.GetRepositories()
	if err != nil {
		return nil, err
	}
	defer repos.Rollback() //nolint

	r, err := repos.Document.List(ctx, &document.ListDocumentArgs{
		PageArgs: document.PageArgs{
			PageSize:  queryParams.PageSize,
			PageToken: queryParams.PageToken,
		},
	})
	if err != nil {
		return nil, err
	}

	ds, err := present.InternalDocumentsToAPIDocuments(r.Documents)
	if err != nil {
		return nil, err
	}

	return &serverdef.ListDocumentsResponse{
		Documents:     ds,
		NextPageToken: r.NextPageToken,
	}, nil
}
