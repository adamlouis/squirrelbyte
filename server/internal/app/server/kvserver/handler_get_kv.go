package kvserver

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/kvmodel"
)

func (h *hdl) GetKV(ctx context.Context, pathParams *kvmodel.GetKVPathParams) (*kvmodel.KV, error) {
	repo, _, rollback, err := h.GetRepository()
	if err != nil {
		return nil, err
	}
	defer rollback() //nolint

	return repo.Get(ctx, pathParams.Key)
}
