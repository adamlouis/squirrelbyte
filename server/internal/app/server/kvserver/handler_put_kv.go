package kvserver

import (
	"context"
	"fmt"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/kvmodel"
)

func (h *hdl) PutKV(ctx context.Context, pathParams *kvmodel.PutKVPathParams, body *kvmodel.KV) (*kvmodel.KV, error) {
	if pathParams.Key != body.Key {
		return nil, fmt.Errorf("key in path does not match key in request body")
	}
	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return nil, err
	}
	defer rollback() //nolint

	out, err := repo.Put(ctx, body)
	if err != nil {
		return nil, err
	}

	if err := commit(); err != nil {
		return nil, err
	}

	return out, nil
}
