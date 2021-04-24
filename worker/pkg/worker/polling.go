package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/adamlouis/squirrelbyte/server/pkg/client"
)

const (
	errorSleepDuration = time.Second * 1
)

func NewPollingWorker(jobClient client.JobClient) Worker {
	return &wkr{
		jobClient: jobClient,
		working:   false,
		fnsByName: map[string]WorkerFn{},
	}
}

type wkr struct {
	jobClient client.JobClient
	working   bool
	fnsByName map[string]WorkerFn
}

func (w *wkr) Register(name string, fn WorkerFn) error {
	if w.working {
		return fmt.Errorf("cannot register a function while the worker is working")
	}

	if _, ok := w.fnsByName[name]; ok {
		return fmt.Errorf("cannot register multiple functions with the name")
	}

	w.fnsByName[name] = fn
	return nil
}

func (w *wkr) Work(ctx context.Context) error {
	w.working = true
	defer func() {
		w.working = false
	}()

	for {
		j, err := w.jobClient.Claim(ctx)
		if err != nil {
			fmt.Println(err) // TODO - errs to chan?
			time.Sleep(errorSleepDuration)
			continue
		}

		fmt.Println("claimed", j.ID, j.Name, w.fnsByName)

		fn := w.fnsByName[j.Name]
		if fn == nil {
			err = w.jobClient.Release(ctx, j.ID)
			if err != nil {
				fmt.Println(j.ID, err) // TODO - errs to chan?
			} else {
				fmt.Println("released", j.ID)
			}
			time.Sleep(errorSleepDuration)
			continue
		}

		go func() {
			err := fn(ctx, j)
			if err != nil {
				err = w.jobClient.SetError(ctx, j.ID, map[string]interface{}{
					"error": fmt.Sprintf("%v", err),
				})
				if err != nil {
					fmt.Println(j.ID, err) // TODO - errs to chan?
				}
			} else {
				noop := map[string]interface{}{} // TODO - rm `output` param ... adds confusion to where results go (wherever, but not in the job itself)
				err = w.jobClient.SetSuccess(ctx, j.ID, noop)
				if err != nil {
					fmt.Println(j.ID, err) // TODO - errs to chan?
				}
			}
		}()
	}
}
