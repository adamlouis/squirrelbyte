package jobserver

import (
	"context"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

func (h *hdl) ReleaseJob(ctx context.Context, pathParams *jobmodel.ReleaseJobPathParams) (*jobmodel.Job, int, error) {
	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer rollback() //nolint

	released, err := repo.Release(ctx, pathParams.JobID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err = commit(); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	out, err := present.InternalJobToAPIJob(released)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return out, http.StatusOK, nil
}
