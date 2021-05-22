package kvserver

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/kvmodel"
)

func (h *hdl) DeleteKV(ctx context.Context, pathParams *kvmodel.DeleteKVPathParams) error {
	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return err
	}
	defer rollback() //nolint

	if err = repo.Delete(ctx, pathParams.Key); err != nil {
		return err
	}

	return commit()
}
