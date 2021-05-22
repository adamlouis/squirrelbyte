package jobserver

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

func (h *hdl) SetJobError(ctx context.Context, pathParams *jobmodel.SetJobErrorPathParams) (*jobmodel.Job, error) {
	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return nil, err
	}
	defer rollback() //nolint

	out, err := repo.Error(ctx, pathParams.JobID)
	if err != nil {
		return nil, err
	}

	if err = commit(); err != nil {
		return nil, err
	}

	return out, nil
}
