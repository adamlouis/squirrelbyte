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

	pollingRunner := worker.NewPollingRunner(jobClient)

	hn := hackernews.Integration{JobClient: jobClient}

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
