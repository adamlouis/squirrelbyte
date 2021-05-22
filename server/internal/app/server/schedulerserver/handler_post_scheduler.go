package schedulerserver

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/schedulermodel"
)

func (h *hdl) PostScheduler(ctx context.Context, body *schedulermodel.Scheduler) (*schedulermodel.Scheduler, error) {
	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return nil, err
	}
	defer rollback()

	out, err := repo.Put(ctx, body)
	if err != nil {
		return nil, err
	}

	if err := commit(); err != nil {
		return nil, err
	}

	h.onUpdate(body.ID)

	return out, nil
}
