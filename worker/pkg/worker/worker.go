package worker

import (
	"context"
)

type WorkerFn func(ctx context.Context, input map[string]interface{}) error

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
