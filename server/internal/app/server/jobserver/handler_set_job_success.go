package jobserver

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

func (h *hdl) SetJobSuccess(ctx context.Context, pathParams *jobmodel.SetJobSuccessPathParams) (*jobmodel.Job, error) {
	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return nil, err
	}
	defer rollback() //nolint

	out, err := repo.Success(ctx, pathParams.JobID)
	if err != nil {
		return nil, err
	}

	if err = commit(); err != nil {
		return nil, err
	}

	return out, nil
}
