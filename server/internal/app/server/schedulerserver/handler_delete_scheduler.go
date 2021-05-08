package schedulerserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/schedulermodel"
)

func (h *hdl) DeleteScheduler(ctx context.Context, pathParams *schedulermodel.DeleteSchedulerPathParams) (int, error) {
	return http.StatusInternalServerError, fmt.Errorf("unimplemented")
}
