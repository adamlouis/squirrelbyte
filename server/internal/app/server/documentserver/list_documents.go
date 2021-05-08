package documentserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
)

func (h *hdl) ListDocuments(ctx context.Context, queryParams *documentmodel.ListDocumentsQueryParams) (*documentmodel.ListDocumentsResponse, int, error) {
	return nil, http.StatusInternalServerError, fmt.Errorf("unimplemented")
}
