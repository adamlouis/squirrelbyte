package serverhandler

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

func (a *apiHandler) DeleteJob(ctx context.Context, pathParams *model.DeleteJobPathParams) error {
	repos, err := a.GetRepositories()
	if err != nil {
		return err
	}
	defer repos.Rollback() //nolint

	err = repos.Job.Delete(ctx, pathParams.JobID)
	if err != nil {
		return err
	}

	if err := repos.Commit(); err != nil {
		return nil
	}

	return repos.Commit()
}
