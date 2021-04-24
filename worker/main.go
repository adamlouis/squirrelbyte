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

// todo
// - provide workers w/ oauth access tokens, 3p credentials, etc.
// - web socket
// - scheduler to queue at times
// - "scheduled for?"
// - integrations: strava, spotify, hacker news, github, aws

func main() {
	ctx := context.Background()

	jobClient := client.NewHTTPJobClient(jobDomain)

	pollingWorker := worker.NewPollingWorker(jobClient)

	err := pollingWorker.Register("hackernews.GetTop", hackernews.GetTop)
	if err != nil {
		log.Fatal(err)
	}

	pollingWorker.Work(ctx)
}
