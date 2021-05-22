package schedulerserver

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/schedulermodel"
)

func (h *hdl) GetScheduler(ctx context.Context, pathParams *schedulermodel.GetSchedulerPathParams) (*schedulermodel.Scheduler, error) {
	repo, _, rollback, err := h.GetRepository()
	if err != nil {
		return nil, err
	}
	defer rollback()

	return repo.Get(ctx, pathParams.SchedulerID)
}
