package jobserver

import (
	"context"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/job"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

func (h *hdl) ListJobs(ctx context.Context, queryParams *jobmodel.ListJobsQueryParams) (*jobmodel.ListJobsResponse, int, error) {
	repo, _, rollback, err := h.GetRepository()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer rollback() //nolint

	listed, err := repo.List(ctx, &job.ListJobArgs{
		PageArgs: job.PageArgs{
			PageSize:  queryParams.PageSize,
			PageToken: queryParams.PageToken,
		},
	})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	js, err := present.InternalJobsToAPIJobs(listed.Jobs)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &jobmodel.ListJobsResponse{
		Jobs:          js,
		NextPageToken: listed.NextPageToken,
	}, http.StatusOK, nil
}
