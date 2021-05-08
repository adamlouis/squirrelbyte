package scheduler

import (
	"context"
	"fmt"

	"github.com/adamlouis/squirrelbyte/server/pkg/client"
	"github.com/adamlouis/squirrelbyte/server/pkg/model"
	"github.com/robfig/cron/v3"
)

// TODO - api & persistence, let caller do

type Scheduler interface {
	Run(ctx context.Context)
}

func NewScheduler(jobClient client.JobClient) Scheduler {
	return &schd{
		jobClient: jobClient,
	}
}

type schd struct {
	jobClient client.JobClient
}

func (sc *schd) Run(ctx context.Context) {
	c := cron.New()

	c.AddFunc("0 * * * *", func() {
		err := sc.jobClient.Queue(ctx, "hackernews.GetTop", model.EmptyJSON())
		fmt.Println("scheduling!", err)
	})

	c.Run()
}
