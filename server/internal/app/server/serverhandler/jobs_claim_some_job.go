package serverhandler

import (
	"context"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/internal/app/server/serverdef"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/job"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

func (a *apiHandler) ClaimJob(ctx context.Context, pathParams *model.ClaimJobPathParams) (*model.Job, error) {
	repos, err := a.GetRepositories()
	if err != nil {
		return nil, err
	}
	defer repos.Rollback() //nolint

	claimed, err := repos.Job.Claim(ctx, job.ClaimOptions{
		JobID: pathParams.JobID,
	})
	if err != nil {
		return nil, err
	}

	if claimed == nil {
		return nil, serverdef.NewHTTPErrorFromString(http.StatusNoContent, "")
	}

	if err := repos.Commit(); err != nil {
		return nil, err
	}

	return present.InternalJobToAPIJob(claimed)
}
