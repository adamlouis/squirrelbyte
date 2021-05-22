package jobserver

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/job"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

func (h *hdl) ClaimSomeJob(ctx context.Context, body *jobmodel.ClaimSomeJobRequest) (*jobmodel.Job, error) {
	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return nil, err
	}
	defer rollback() //nolint

	out, err := repo.Claim(ctx, &job.ClaimOptions{
		Names: body.Names,
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
