package schedulerserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/schedulermodel"
)

func (h *hdl) GetScheduler(ctx context.Context, pathParams *schedulermodel.GetSchedulerPathParams) (*schedulermodel.Scheduler, int, error) {
	return nil, http.StatusInternalServerError, fmt.Errorf("unimplemented")
}
