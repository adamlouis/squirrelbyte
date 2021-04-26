package hackernews

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/client"
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
// - per-api limiting?
// - per-job name rate limiting?
// - total job rate limiting?
// - scheduler to queue jobs on cron, or whatever

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

func (i *Integration) getTopStoriesFn(ctx context.Context, input map[string]interface{}) error {
	ids, err := fetchItemIDs("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		return err
	}

	for _, id := range ids {
		err = i.JobClient.Queue(ctx, "hackernews.GetItem", map[string]interface{}{"item_id": id})
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

func (i *Integration) getItemFn(ctx context.Context, input map[string]interface{}) error {
	b, err := json.Marshal(input)
	if err != nil {
		return err
	}

	inputStruct := GetItemInput{}
	err = json.Unmarshal(b, &input)
	if err != nil {
		return err
	}

	item, err := fetchItem(inputStruct.ItemID)
	if err != nil {
		return err
	}

	// TODO: write to server - document client
	fmt.Println(string(item))

	return nil
}

// todo : pattern for rate-limiting
func fetchItem(id int) ([]byte, error) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
	resp, err := http.Get(url) //nolint
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func fetchItemIDs(url string) ([]int, error) {
	resp, err := http.Get(url) //nolint
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ids := []int{}
	if err := json.NewDecoder(resp.Body).Decode(&ids); err != nil {
		return nil, err
	}

	return ids, nil
}
