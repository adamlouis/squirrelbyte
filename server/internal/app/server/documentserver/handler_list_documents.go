package documentserver

import (
	"context"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/document"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
)

func (h *hdl) ListDocuments(ctx context.Context, queryParams *documentmodel.ListDocumentsQueryParams) (*documentmodel.ListDocumentsResponse, int, error) {
	repo, _, rollback, err := h.GetRepository()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer rollback() //nolint

	r, err := repo.List(ctx, &document.ListDocumentArgs{
		PageArgs: document.PageArgs{
			PageSize:  queryParams.PageSize,
			PageToken: queryParams.PageToken,
		},
	})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	ds, err := present.InternalDocumentsToAPIDocuments(r.Documents)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &documentmodel.ListDocumentsResponse{
		Documents:     ds,
		NextPageToken: r.NextPageToken,
	}, http.StatusOK, nil
}
