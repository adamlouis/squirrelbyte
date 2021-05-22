package jobserver

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

func (h *hdl) DeleteJob(ctx context.Context, pathParams *jobmodel.DeleteJobPathParams) error {
	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return err
	}
	defer rollback() //nolint

	err = repo.Delete(ctx, pathParams.JobID)
	if err != nil {
		return err
	}

	return commit()
}
