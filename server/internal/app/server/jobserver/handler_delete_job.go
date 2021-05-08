package jobserver

import (
	"context"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

func (h *hdl) DeleteJob(ctx context.Context, pathParams *jobmodel.DeleteJobPathParams) (int, error) {
	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer rollback() //nolint

	err = repo.Delete(ctx, pathParams.JobID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if err = commit(); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
