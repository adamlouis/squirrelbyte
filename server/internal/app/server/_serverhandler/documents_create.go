package serverhandler

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

func (a *apiHandler) CreateDocument(ctx context.Context, document *model.Document) (*model.Document, error) {
	return a.PutDocument(ctx, &model.PutDocumentPathParams{
		DocumentID: document.ID,
	}, document)
}
