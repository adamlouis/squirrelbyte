package worker

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/adamlouis/squirrelbyte/server/pkg/client/jobclient"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

const (
	errorSleepDuration = time.Second * 1
)

func NewPollingRunner(jobClient jobclient.Client, maxConcurrent int) Runner {
	return &rnr{
		maxConcurrent: maxConcurrent,
		jobClient:     jobClient,
		running:       false,
		fnsByName:     map[string]WorkerFn{},
	}
}

type rnr struct {
	maxConcurrent int
	jobClient     jobclient.Client
	running       bool
	fnsByName     map[string]WorkerFn
}

func (r *rnr) Register(w *Worker) error {
	if r.running {
		return fmt.Errorf("cannot register a worker function while the runner is running")
	}

	if _, ok := r.fnsByName[w.Name]; ok {
		return fmt.Errorf("cannot register multiple functions with the name")
	}

	r.fnsByName[w.Name] = w.Fn

	fmt.Println("registered", w.Name)

	return nil
}

func (r *rnr) Run(ctx context.Context) error {
	r.running = true
	defer func() {
		r.running = false
	}()

	jobNames := make([]string, len(r.fnsByName))
	i := 0
	for n := range r.fnsByName {
		jobNames[i] = n
		i++
	}

	jobChan := make(chan struct{}, r.maxConcurrent)

	for {
		j, s, err := r.jobClient.ClaimSomeJob(ctx, &jobmodel.ClaimSomeJobRequest{
			Names: jobNames,
		})
		if s == http.StatusNoContent {
			fmt.Println("no jobs queued") // TODO - errs to chan?
			time.Sleep(errorSleepDuration)
			continue
		}
		if err != nil {
			fmt.Println(err) // TODO - errs to chan?
			time.Sleep(errorSleepDuration)
			continue
		}

		fmt.Println("claimed", j, s, j.ID, j.Name)

		fn := r.fnsByName[j.Name]
		if fn == nil {
			return fmt.Errorf("no handler registered for job %s", j.Name) // TODO - close channels?
		}

		jobChan <- struct{}{}
		go func(c <-chan struct{}) {
			// TODO - rm print statements, trace starting here
			r.do(ctx, fn, j)
			<-c
		}(jobChan)
	}
}

func (r *rnr) do(ctx context.Context, fn WorkerFn, j *jobmodel.Job) {
	err := fn(ctx, j.Input)
	if err != nil {
		fmt.Println(err) // TODO: output
		_, _, err = r.jobClient.SetJobError(ctx, &jobmodel.SetJobErrorPathParams{JobID: j.ID})
		if err != nil {
			fmt.Println(j.ID, err) // TODO - errs to chan?
		} else {
			fmt.Println("set error", j.ID)
		}
	} else {
		_, _, err = r.jobClient.SetJobSuccess(ctx, &jobmodel.SetJobSuccessPathParams{JobID: j.ID})
		if err != nil {
			fmt.Println(j.ID, err) // TODO - errs to chan?
		} else {
			fmt.Println("set success", j.ID)
		}
	}
}
