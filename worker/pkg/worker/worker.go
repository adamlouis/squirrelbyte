package worker

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

type WorkerFn func(ctx context.Context, j *model.Job) error

type Worker interface {
	Register(name string, fn WorkerFn) error
	Work(context.Context) error
}
