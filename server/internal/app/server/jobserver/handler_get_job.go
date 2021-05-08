package jobserver

import (
	"context"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

func (h *hdl) GetJob(ctx context.Context, pathParams *jobmodel.GetJobPathParams) (*jobmodel.Job, int, error) {
	repo, _, rollback, err := h.GetRepository()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer rollback() //nolint

	got, err := repo.Get(ctx, pathParams.JobID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	out, err := present.InternalJobToAPIJob(got)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return out, http.StatusOK, nil
}
