package schedulerserver

import (
	"context"
	"fmt"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/schedulermodel"
)

func (h *hdl) PutScheduler(ctx context.Context, pathParams *schedulermodel.PutSchedulerPathParams, body *schedulermodel.Scheduler) (*schedulermodel.Scheduler, error) {
	if pathParams.SchedulerID != body.ID {
		return nil, fmt.Errorf("id in path does not match id in request body")
	}
	return h.PostScheduler(ctx, body)
}
