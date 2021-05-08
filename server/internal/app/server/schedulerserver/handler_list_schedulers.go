package schedulerserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/schedulermodel"
)

func (h *hdl) ListSchedulers(ctx context.Context, queryParams *schedulermodel.ListSchedulersRequest) (*schedulermodel.ListSchedulersResponse, int, error) {
	return nil, http.StatusInternalServerError, fmt.Errorf("unimplemented")
}
