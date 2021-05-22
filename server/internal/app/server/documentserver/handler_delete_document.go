package documentserver

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
)

func (h *hdl) DeleteDocument(ctx context.Context, pathParams *documentmodel.DeleteDocumentPathParams) error {
	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return err
	}
	defer rollback() //nolint

	err = repo.Delete(ctx, pathParams.DocumentID)
	if err != nil {
		return err
	}

	return commit()
}
