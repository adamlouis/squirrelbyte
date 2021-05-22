package jobserver

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/job"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

func (h *hdl) ClaimJob(ctx context.Context, pathParams *jobmodel.ClaimJobPathParams) (*jobmodel.Job, error) {
	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return nil, err
	}
	defer rollback() //nolint

	out, err := repo.Claim(ctx, &job.ClaimOptions{
		JobID: pathParams.JobID,
	})
	if err != nil {
		return nil, err
	}

	if out == nil {
		return nil, nil
	}

	if err = commit(); err != nil {
		return nil, err
	}

	return out, nil
}
