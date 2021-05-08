package documentserver

import (
	"context"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
)

func (h *hdl) DeleteDocument(ctx context.Context, pathParams *documentmodel.DeleteDocumentPathParams) (int, error) {
	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer rollback() //nolint

	err = repo.Delete(ctx, pathParams.DocumentID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = commit()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
