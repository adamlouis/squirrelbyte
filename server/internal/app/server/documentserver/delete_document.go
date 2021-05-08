package documentserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
)

func (h *hdl) DeleteDocument(ctx context.Context, pathParams *documentmodel.DeleteDocumentPathParams) (int, error) {
	return http.StatusInternalServerError, fmt.Errorf("unimplemented")
}
