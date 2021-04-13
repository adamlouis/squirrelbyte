package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/adamlouis/squirrelbyte/server/internal/app/server/serverdef"
)

func respToMap(r *http.Response) (map[string]interface{}, error) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var itemRespBodyMap map[string]interface{}
	err = json.Unmarshal(bytes, &itemRespBodyMap)
	if err != nil {
		return nil, err
	}
	return itemRespBodyMap, nil
}

func ifcToBuf(ifc interface{}) (*bytes.Buffer, error) {
	bs, err := json.Marshal(ifc)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(bs), nil
}

func loadItems(ids []int) error {
	for _, id := range ids {
		itemURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)

		itemResp, err := http.Get(itemURL) //nolint
		if err != nil {
			return err
		}

		itemMap, err := respToMap(itemResp)
		if err != nil {
			return err
		}

		reqBuf, err := ifcToBuf(&serverdef.Document{
			ID: fmt.Sprintf("hn.item.%d", id),
			Header: map[string]interface{}{
				"url": itemURL,
			},
			Body: itemMap,
		})
		if err != nil {
			return err
		}

		createResp, err := http.Post("http://localhost:9922/api/documents", "application/json", reqBuf)
		if err != nil {
			return err
		}

		createRespBytes, err := ioutil.ReadAll(createResp.Body)
		if err != nil {
			return err
		}

		fmt.Print(string(createRespBytes))
		time.Sleep(3 * time.Second)
	}
	return nil
}

func loadList(url string) error {
	resp, err := http.Get(url) //nolint
	if err != nil {
		return err
	}

	ids := []int{}
	if err := json.NewDecoder(resp.Body).Decode(&ids); err != nil {
		return err
	}

	return loadItems(ids)
}

func main() {
	_, err := http.Get("http://localhost:9922/api/status")
	if err != nil {
		log.Fatal(err)
	}

	urls := []string{
		"https://hacker-news.firebaseio.com/v0/topstories.json",
		"https://hacker-news.firebaseio.com/v0/newstories.json",
		"https://hacker-news.firebaseio.com/v0/askstories.json",
		"https://hacker-news.firebaseio.com/v0/showstories.json",
	}

	for _, url := range urls {
		err = loadList(url)
		if err != nil {
			log.Fatal(err)
		}
	}
}
