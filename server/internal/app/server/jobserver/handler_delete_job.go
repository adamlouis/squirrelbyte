package jobserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

func (h *hdl) DeleteJob(ctx context.Context, pathParams *jobmodel.DeleteJobPathParams) (int, error) {
	return http.StatusInternalServerError, fmt.Errorf("unimplemented")
}
