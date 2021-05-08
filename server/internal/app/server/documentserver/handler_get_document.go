package documentserver

import (
	"context"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
)

func (h *hdl) GetDocument(ctx context.Context, pathParams *documentmodel.GetDocumentPathParams) (*documentmodel.Document, int, error) {
	repo, _, rollback, err := h.GetRepository()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer rollback() //nolint

	d, err := repo.Get(ctx, pathParams.DocumentID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	out, err := present.InternalDocumentToAPIDocument(d)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return out, http.StatusOK, nil
}
