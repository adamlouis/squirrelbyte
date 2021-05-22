package jobserver

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

func (h *hdl) GetJob(ctx context.Context, pathParams *jobmodel.GetJobPathParams) (*jobmodel.Job, error) {
	repo, _, rollback, err := h.GetRepository()
	if err != nil {
		return nil, err
	}
	defer rollback() //nolint

	return repo.Get(ctx, pathParams.JobID)
}
