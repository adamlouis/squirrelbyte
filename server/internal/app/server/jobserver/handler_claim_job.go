package jobserver

import (
	"context"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/job"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

func (h *hdl) ClaimJob(ctx context.Context, pathParams *jobmodel.ClaimJobPathParams) (*jobmodel.Job, int, error) {
	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer rollback() //nolint

	claimed, err := repo.Claim(ctx, job.ClaimOptions{
		JobID: pathParams.JobID,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if claimed == nil {
		return nil, http.StatusNoContent, nil
	}

	if err = commit(); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	out, err := present.InternalJobToAPIJob(claimed)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return out, http.StatusOK, nil
}
