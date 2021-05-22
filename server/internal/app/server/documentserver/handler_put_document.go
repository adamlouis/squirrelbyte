package documentserver

import (
	"context"
	"fmt"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
)

func (h *hdl) PutDocument(ctx context.Context, pathParams *documentmodel.PutDocumentPathParams, body *documentmodel.Document) (*documentmodel.Document, error) {
	if pathParams.DocumentID != body.ID {
		return nil, fmt.Errorf("id in path does not match id in request body")
	}
	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return nil, err
	}
	defer rollback() //nolint

	out, err := repo.Put(ctx, body)
	if err != nil {
		return nil, err
	}

	if err := commit(); err != nil {
		return nil, err
	}

	return out, nil
}
