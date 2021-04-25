package hackernews

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/client"
	"github.com/adamlouis/squirrelbyte/server/pkg/model"
	"github.com/adamlouis/squirrelbyte/worker/pkg/worker"
)

// todo:
// X add claimed_at
// X only poll jobs matching names
// - handle concurrency gracefully server-side
// - better interfaces for integration
// - trigger job from within worker
// - write results to document server
// - remove output field ... think on it ... optional?
// - wss rather than poll

type Integration struct {
	JobClient client.JobClient
}

type GetItemInput struct {
	ItemID int `json:"item_id"`
}

// TODO: better interface?
func (i *Integration) GetTopStoriesWorker() *worker.Worker {
	return &worker.Worker{
		Name: "hackernews.GetTop",
		Fn:   i.getTopStoriesFn,
	}
}

func (i *Integration) GetItemWorker() *worker.Worker {
	return &worker.Worker{
		Name: "hackernews.GetItem",
		Fn:   i.getItemFn,
	}
}

func (i *Integration) getTopStoriesFn(ctx context.Context, j *model.Job) error {
	ids, err := fetchItemIDs("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		return err
	}

	fmt.Println(ids)

	for _, id := range ids {
		// TODO: better interface
		err = i.JobClient.Queue(ctx, "hackernews.GetItem", map[string]interface{}{"item_id": id})
		fmt.Println(err)
	}

	return nil
}

func (i *Integration) getItemFn(ctx context.Context, j *model.Job) error {
	fmt.Println(j.Input)
	// ids, err := fetchItemIDs("https://hacker-news.firebaseio.com/v0/topstories.json")
	// if err != nil {
	// 	return err
	// }

	// fmt.Println(ids)

	// for _, id := range ids {
	// 	i.jobClient.Queue(ctx, "hackernews.GetItem", map[string]interface{}{"item_id": id})
	// }

	return nil
}

func fetchItemIDs(url string) ([]int, error) {
	resp, err := http.Get(url) //nolint
	if err != nil {
		return nil, err
	}

	ids := []int{}
	if err := json.NewDecoder(resp.Body).Decode(&ids); err != nil {
		return nil, err
	}

	return ids, nil
}
