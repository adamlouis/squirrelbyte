package documentserver

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
)

func (h *hdl) PostDocument(ctx context.Context, body *documentmodel.Document) (*documentmodel.Document, error) {
	return h.PutDocument(ctx, &documentmodel.PutDocumentPathParams{DocumentID: body.ID}, body)
}
