package documentserver

import (
	"context"
	"net/http"
	"time"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/document"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
)

func (h *hdl) QueryDocuments(ctx context.Context, body *documentmodel.QueryDocumentsRequest) (*documentmodel.QueryDocumentsResponse, int, error) {
	start := time.Now()
	repo, _, rollback, err := h.GetRepository()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer rollback() //nolint

	r, err := repo.Query(ctx, &document.Query{
		Select:    body.Select,
		Where:     body.Where,
		GroupBy:   body.GroupBy,
		OrderBy:   body.OrderBy,
		Limit:     body.Limit,
		PageToken: body.PageToken,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &documentmodel.QueryDocumentsResponse{
		Result:        r.Result,
		NextPageToken: r.NextPageToken,
		Insights: map[string]interface{}{
			"duration_ms": time.Since(start) / time.Millisecond,
		},
	}, http.StatusOK, nil

}
