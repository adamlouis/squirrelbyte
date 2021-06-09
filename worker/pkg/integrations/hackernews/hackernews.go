package hackernews

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/client/documentclient"
	"github.com/adamlouis/squirrelbyte/server/pkg/client/jobclient"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
	"github.com/adamlouis/squirrelbyte/worker/pkg/worker"
)

type Integration struct {
	JobClient      jobclient.Client
	DocumentClient documentclient.Client
}

type GetItemInput struct {
	ItemID uint64 `json:"item_id"`
}

// TODO: this interface needs help ... find the pattern
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
		_, _, err = i.JobClient.QueueJob(ctx, &jobmodel.Job{
			Name:  "hackernews.GetItem",
			Input: map[string]interface{}{"item_id": id},
		})
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
	err = json.Unmarshal(b, &inputStruct)
	if err != nil {
		return err
	}

	item, err := fetchItem(inputStruct.ItemID)
	if err != nil {
		return err
	}

	body := map[string]interface{}{}
	err = json.Unmarshal(item, &body)
	if err != nil {
		return err
	}

	_, _, err = i.DocumentClient.PostDocument(ctx, &documentmodel.Document{
		ID: fmt.Sprintf("hn.item.%d", inputStruct.ItemID),
		Header: map[string]interface{}{
			"api_url": getItemURL(inputStruct.ItemID),
			"hn_url":  fmt.Sprintf("https://news.ycombinator.com/item?id=%d", inputStruct.ItemID),
		},
		Body: body,
	})

	return err
}

func getItemURL(id uint64) string {
	return fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
}

func fetchItem(id uint64) ([]byte, error) {
	resp, err := http.Get(getItemURL(id)) //nolint
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
