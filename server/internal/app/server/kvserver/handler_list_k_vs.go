package kvserver

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/kvmodel"
)

func (h *hdl) ListKVs(ctx context.Context, queryParams *kvmodel.ListKVsRequest) (*kvmodel.ListKVsResponse, error) {
	repo, _, rollback, err := h.GetRepository()
	if err != nil {
		return nil, err
	}
	defer rollback() //nolint

	return repo.List(ctx, queryParams)
}
