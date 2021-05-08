package documentserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
)

func (h *hdl) PutDocument(ctx context.Context, pathParams *documentmodel.PutDocumentPathParams, body *documentmodel.Document) (*documentmodel.Document, int, error) {
	return nil, http.StatusInternalServerError, fmt.Errorf("unimplemented")
}
