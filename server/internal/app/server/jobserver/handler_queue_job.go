package jobserver

import (
	"context"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

func (h *hdl) QueueJob(ctx context.Context, requestBody *jobmodel.Job) (*jobmodel.Job, int, error) {
	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer rollback() //nolint

	in, err := present.APIJobToInternalJob(requestBody)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	queued, err := repo.Queue(ctx, in)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err = commit(); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	out, err := present.InternalJobToAPIJob(queued)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return out, http.StatusOK, nil
}
