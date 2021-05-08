package jobserver

import (
	"context"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

func (h *hdl) SetJobSuccess(ctx context.Context, pathParams *jobmodel.SetJobSuccessPathParams) (*jobmodel.Job, int, error) {
	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer rollback() //nolint

	succeeded, err := repo.Success(ctx, pathParams.JobID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err = commit(); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	out, err := present.InternalJobToAPIJob(succeeded)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return out, http.StatusOK, nil
}
