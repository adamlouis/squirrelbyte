package schedulerserver

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/schedulermodel"
)

func (h *hdl) DeleteScheduler(ctx context.Context, pathParams *schedulermodel.DeleteSchedulerPathParams) error {
	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return err
	}
	defer rollback()

	err = repo.Delete(ctx, pathParams.SchedulerID)
	if err != nil {
		return err
	}

	return commit()
}
