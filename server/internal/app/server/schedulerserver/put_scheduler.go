package schedulerserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/schedulermodel"
)

func (h *hdl) PutScheduler(ctx context.Context, pathParams *schedulermodel.PutSchedulerPathParams, body *schedulermodel.Scheduler) (*schedulermodel.Scheduler, int, error) {
	return nil, http.StatusInternalServerError, fmt.Errorf("unimplemented")
}
