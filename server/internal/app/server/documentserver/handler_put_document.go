package documentserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
)

func (h *hdl) PutDocument(ctx context.Context, pathParams *documentmodel.PutDocumentPathParams, body *documentmodel.Document) (*documentmodel.Document, int, error) {
	if pathParams.DocumentID != body.ID {
		return nil, http.StatusBadRequest, fmt.Errorf("document id in path does not match document id in request body")
	}
	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer rollback() //nolint

	in, err := present.APIDocumentToInternalDocument(body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	put, err := repo.Put(ctx, in)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err := commit(); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	out, err := present.InternalDocumentToAPIDocument(put)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return out, http.StatusOK, nil
}
