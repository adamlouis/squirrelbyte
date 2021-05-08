package jobserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

func (h *hdl) QueueJob(ctx context.Context) (*jobmodel.Job, int, error) {
	return nil, http.StatusInternalServerError, fmt.Errorf("unimplemented")
}
