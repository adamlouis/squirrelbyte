package worker

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

type WorkerFn func(ctx context.Context, j *model.Job) error

type Worker struct {
	Name string
	Fn   WorkerFn
}

type Runner interface {
	Register(w *Worker) error
	Run(context.Context) error
}

type Queuer interface {
	Queue(ctx context.Context, name string, input map[string]interface{}) error
}
