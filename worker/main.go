package main

import (
	"context"
	"log"

	"github.com/adamlouis/squirrelbyte/server/pkg/client/documentclient"
	"github.com/adamlouis/squirrelbyte/server/pkg/client/jobclient"
	"github.com/adamlouis/squirrelbyte/server/pkg/client/kvclient"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
	"github.com/adamlouis/squirrelbyte/worker/pkg/integrations/example"
	"github.com/adamlouis/squirrelbyte/worker/pkg/integrations/hackernews"
	"github.com/adamlouis/squirrelbyte/worker/pkg/integrations/spotify"
	"github.com/adamlouis/squirrelbyte/worker/pkg/worker"
)

const (
	jobURL      = "http://localhost:9922/api"
	documentURL = "http://localhost:9922/api"
	kvURL       = "http://localhost:9922/api"
)

type WorkerFn func(ctx context.Context, j *jobmodel.Job) (map[string]interface{}, error)

func main() {
	ctx := context.Background()

	jobClient := jobclient.NewHTTPClient(jobURL)
	documentClient := documentclient.NewHTTPClient(documentURL)
	kvClient := kvclient.NewHTTPClient(kvURL)

	pollingRunner := worker.NewPollingRunner(jobClient, 3)

	hn := hackernews.Integration{JobClient: jobClient, DocumentClient: documentClient}
	ex := example.Integration{JobClient: jobClient, DocumentClient: documentClient}
	sp := spotify.Integration{JobClient: jobClient, DocumentClient: documentClient, KVClient: kvClient}

	workers := []*worker.Worker{
		hn.GetTopStoriesWorker(),
		hn.GetItemWorker(),
		ex.GetTheJobWorker(),
		sp.GetFetchRecentPlaysWorker(),
	}

	for _, w := range workers {
		err := pollingRunner.Register(w)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Fatal(pollingRunner.Run(ctx))
}
