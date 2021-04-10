package serverhandler

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/internal/app/server/serverdef"
)

func (a *apiHandler) CreateDocument(ctx context.Context, document *serverdef.Document) (*serverdef.Document, error) {
	return a.PutDocument(ctx, &serverdef.PutDocumentPathParams{
		DocumentID: document.ID,
	}, document)
}
