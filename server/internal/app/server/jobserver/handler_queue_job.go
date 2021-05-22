package jobserver

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

func (h *hdl) QueueJob(ctx context.Context, requestBody *jobmodel.Job) (*jobmodel.Job, error) {
	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return nil, err
	}
	defer rollback() //nolint

	out, err := repo.Queue(ctx, requestBody)
	if err != nil {
		return nil, err
	}

	if err = commit(); err != nil {
		return nil, err
	}

	return out, nil
}
