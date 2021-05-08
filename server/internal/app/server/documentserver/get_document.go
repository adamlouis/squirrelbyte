package documentserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
)

func (h *hdl) GetDocument(ctx context.Context, pathParams *documentmodel.GetDocumentPathParams) (*documentmodel.Document, int, error) {
	return nil, http.StatusInternalServerError, fmt.Errorf("unimplemented")
}
