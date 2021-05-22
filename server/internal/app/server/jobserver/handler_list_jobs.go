package jobserver

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

func (h *hdl) ListJobs(ctx context.Context, queryParams *jobmodel.ListJobsQueryParams) (*jobmodel.ListJobsResponse, error) {
	repo, _, rollback, err := h.GetRepository()
	if err != nil {
		return nil, err
	}
	defer rollback() //nolint

	return repo.List(ctx, queryParams)
}
