package example

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/client/documentclient"
	"github.com/adamlouis/squirrelbyte/server/pkg/client/jobclient"
	"github.com/adamlouis/squirrelbyte/worker/pkg/worker"
)

type Integration struct {
	JobClient      jobclient.Client
	DocumentClient documentclient.Client
}

// TODO: this interface needs help ... find the pattern
func (i *Integration) GetTheJobWorker() *worker.Worker {
	return &worker.Worker{
		Name: "TheJob",
		Fn:   i.doTheJob,
	}
}

func (i *Integration) doTheJob(ctx context.Context, input map[string]interface{}) error {
	return nil
}
