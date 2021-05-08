package documentserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
)

func (h *hdl) QueryDocuments(ctx context.Context, body *documentmodel.QueryDocumentsRequest) (*documentmodel.QueryDocumentsResponse, int, error) {
	return nil, http.StatusInternalServerError, fmt.Errorf("unimplemented")
}
