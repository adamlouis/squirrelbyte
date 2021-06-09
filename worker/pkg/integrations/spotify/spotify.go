package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/adamlouis/squirrelbyte/server/pkg/client/documentclient"
	"github.com/adamlouis/squirrelbyte/server/pkg/client/jobclient"
	"github.com/adamlouis/squirrelbyte/server/pkg/client/kvclient"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/kvmodel"
	"github.com/adamlouis/squirrelbyte/worker/pkg/worker"
	"golang.org/x/oauth2"
)

type Integration struct {
	JobClient      jobclient.Client
	DocumentClient documentclient.Client
	KVClient       kvclient.Client
}

// TODO: this interface needs help ... find the pattern
func (i *Integration) GetFetchRecentPlaysWorker() *worker.Worker {
	return &worker.Worker{
		Name: "spotify.FetchRecentPlays",
		Fn:   i.fetchRecentPlays,
	}
}

type FetchRecentPlaysInput struct {
	OAuthKey string `json:"oauth_key"`
}

type RecentlyPlayedResponse struct {
	Items []interface{} `json:"items"`
}

type RecentPlayItem struct {
	PlayedAt string `json:"played_at"`
}

func (i *Integration) fetchRecentPlays(ctx context.Context, input map[string]interface{}) error {
	b, err := json.Marshal(input)
	if err != nil {
		return err
	}

	inputStruct := FetchRecentPlaysInput{}
	err = json.Unmarshal(b, &inputStruct)
	if err != nil {
		return err
	}

	kv, _, err := i.KVClient.GetKV(ctx, &kvmodel.GetKVPathParams{Key: inputStruct.OAuthKey})
	if err != nil {
		return err
	}

	fmt.Println("INPUT -> KV:", inputStruct, kv)

	jb, err := json.Marshal(kv.Value)
	if err != nil {
		return err
	}

	tk := oauth2.Token{}
	json.Unmarshal(jb, &tk)

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/player/recently-played", nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+tk.AccessToken)
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	resb, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(resb))

	ress := RecentlyPlayedResponse{}
	json.Unmarshal(resb, &ress)

	for _, item := range ress.Items {
		ib, _ := json.Marshal(item)
		rpi := RecentPlayItem{}
		json.Unmarshal(ib, &rpi)
		id := rpi.PlayedAt
		id = strings.ReplaceAll(id, ".", "")
		id = strings.ReplaceAll(id, ":", "")
		id = strings.ReplaceAll(id, "-", "")
		body := item.(map[string]interface{})
		header := map[string]interface{}{}
		d, s, e := i.DocumentClient.PutDocument(
			ctx,
			&documentmodel.PutDocumentPathParams{
				DocumentID: id,
			},
			&documentmodel.Document{
				ID:     id,
				Body:   body,
				Header: header,
			})
		fmt.Println(id, d, s, e)
	}
	// fmt.Println(ress)
	// res, err := http.Get("https://api.spotify.com/v1/me/player/recently-played")
	// fmt.Println(res, err)

	return fmt.Errorf("unimplemented")
}
