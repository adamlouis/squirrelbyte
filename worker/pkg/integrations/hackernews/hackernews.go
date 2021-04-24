package hackernews

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model"
	"github.com/adamlouis/squirrelbyte/worker/pkg/worker"
)

// todo:
// - remove output field
// - add claimed_at
// - trigger job from within worker
// - only poll jobs matching names
// - write results to document server
// - wss rather than poll

func GetItem(ctx context.Context, j *model.Job) error {
	return nil
}

func GetTop(ctx context.Context, j *model.Job) error {
	ids, err := fetchItemIDs("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		return err
	}

	// for _, id := range ids {
	// 	// GetItem(ctx, id) // TODO: figure out nice way to queue job from in a worker fn
	// }

	fmt.Println(ids)

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

var (
	_ worker.WorkerFn = GetItem
)
