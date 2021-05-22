package documentserver

import (
	"context"
	"time"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
)

func (h *hdl) QueryDocuments(ctx context.Context, body *documentmodel.QueryDocumentsRequest) (*documentmodel.QueryDocumentsResponse, error) {
	start := time.Now()
	repo, _, rollback, err := h.GetRepository()
	if err != nil {
		return nil, err
	}
	defer rollback() //nolint

	out, err := repo.Query(ctx, body)
	if err != nil {
		return nil, err
	}

	out.Insights = map[string]interface{}{
		"duration_ms": time.Since(start) / time.Millisecond,
	}

	return out, nil
}
