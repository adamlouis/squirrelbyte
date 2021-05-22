package scheduler

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/schedulermodel"
)

type Repository interface {
	Put(ctx context.Context, scheduler *schedulermodel.Scheduler) (*schedulermodel.Scheduler, error)
	Get(ctx context.Context, id string) (*schedulermodel.Scheduler, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, body *schedulermodel.ListSchedulersRequest) (*schedulermodel.ListSchedulersResponse, error)
}

type Runner interface {
	Run(ctx context.Context) error
	Update(ctx context.Context, id string)
}
