package serverhandler

import (
	"context"
	"encoding/json"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

func (a *apiHandler) SetJobSuccess(ctx context.Context, pathParams *model.SetJobSuccessPathParams, job *model.Job) (*model.Job, error) {
	repos, err := a.GetRepositories()
	if err != nil {
		return nil, err
	}
	defer repos.Rollback() //nolint

	output, err := json.Marshal(job.Output)
	if err != nil {
		return nil, err
	}

	got, err := repos.Job.Success(ctx, pathParams.JobID, output)
	if err != nil {
		return nil, err
	}

	if err := repos.Commit(); err != nil {
		return nil, err
	}

	return present.InternalJobToAPIJob(got)
}
