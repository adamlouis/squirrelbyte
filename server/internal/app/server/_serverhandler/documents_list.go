package serverhandler

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/document"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

func (a *apiHandler) ListDocuments(ctx context.Context, queryParams *model.ListDocumentsQueryParams) (*model.ListDocumentsResponse, error) {
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

	return &model.ListDocumentsResponse{
		Documents:     ds,
		NextPageToken: r.NextPageToken,
	}, nil
}
