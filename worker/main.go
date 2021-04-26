package main

import (
	"context"
	"log"

	"github.com/adamlouis/squirrelbyte/server/pkg/client"
	"github.com/adamlouis/squirrelbyte/server/pkg/model"
	"github.com/adamlouis/squirrelbyte/worker/pkg/integrations/hackernews"
	"github.com/adamlouis/squirrelbyte/worker/pkg/worker"
)

const (
	jobDomain = "http://localhost:9922"
)

type WorkerFn func(ctx context.Context, j *model.Job) (map[string]interface{}, error)

func main() {
	ctx := context.Background()

	jobClient := client.NewHTTPJobClient(jobDomain)
	documentClient := client.NewHTTPDocumentClient(jobDomain)

	pollingRunner := worker.NewPollingRunner(jobClient, 3)

	hn := hackernews.Integration{JobClient: jobClient, DocumentClient: documentClient}

	err := pollingRunner.Register(hn.GetTopStoriesWorker())
	if err != nil {
		log.Fatal(err)
	}

	err = pollingRunner.Register(hn.GetItemWorker())
	if err != nil {
		log.Fatal(err)
	}

	pollingRunner.Run(ctx)
}
