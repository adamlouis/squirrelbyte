package jobserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

func (h *hdl) GetJob(ctx context.Context, pathParams *jobmodel.GetJobPathParams) (*jobmodel.Job, int, error) {
	return nil, http.StatusInternalServerError, fmt.Errorf("unimplemented")
}
