package serverhandler

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

func (a *apiHandler) SetJobError(ctx context.Context, pathParams *model.SetJobErrorPathParams, job *model.Job) (*model.Job, error) {
	repos, err := a.GetRepositories()
	if err != nil {
		return nil, err
	}
	defer repos.Rollback() //nolint

	updated, err := repos.Job.Error(ctx, pathParams.JobID)
	if err != nil {
		return nil, err
	}

	if err := repos.Commit(); err != nil {
		return nil, err
	}

	return present.InternalJobToAPIJob(updated)
}
