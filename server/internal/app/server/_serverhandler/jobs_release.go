package serverhandler

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

func (a *apiHandler) ReleaseJob(ctx context.Context, pathParams *model.ReleaseJobPathParams) (*model.Job, error) {
	repos, err := a.GetRepositories()
	if err != nil {
		return nil, err
	}
	defer repos.Rollback() //nolint

	got, err := repos.Job.Release(ctx, pathParams.JobID)
	if err != nil {
		return nil, err
	}

	if err := repos.Commit(); err != nil {
		return nil, err
	}

	return present.InternalJobToAPIJob(got)
}
