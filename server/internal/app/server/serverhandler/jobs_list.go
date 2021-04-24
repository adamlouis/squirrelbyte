package serverhandler

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/job"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

func (a *apiHandler) ListJobs(ctx context.Context, queryParams *model.ListJobsQueryParams) (*model.ListJobsResponse, error) {
	repos, err := a.GetRepositories()
	if err != nil {
		return nil, err
	}
	defer repos.Rollback() //nolint

	listed, err := repos.Job.List(ctx, &job.ListJobArgs{
		PageArgs: job.PageArgs{
			PageSize:  queryParams.PageSize,
			PageToken: queryParams.PageToken,
		},
	})
	if err != nil {
		return nil, err
	}

	js, err := present.InternalJobsToAPIJobs(listed.Jobs)
	if err != nil {
		return nil, err
	}

	return &model.ListJobsResponse{
		Jobs:          js,
		NextPageToken: listed.NextPageToken,
	}, nil
}
