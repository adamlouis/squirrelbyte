package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

type DocumentClient interface {
	Post(ctx context.Context, d *model.Document) (*model.Document, error)
}

func NewHTTPDocumentClient(url string) DocumentClient {
	return &documentClient{
		url: url,
	}
}

type documentClient struct {
	url string
}

func (dc *documentClient) Post(ctx context.Context, d *model.Document) (*model.Document, error) {
	b, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(
		fmt.Sprintf("%s/api/documents", dc.url),
		"application/json",
		bytes.NewBuffer(b),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		b, _ := ioutil.ReadAll(res.Body)
		return nil, fmt.Errorf("error posting document: %s: %s", res.Status, string(b))
	}

	doc := model.Document{}
	err = json.NewDecoder(res.Body).Decode(&doc)
	if err != nil {
		return nil, fmt.Errorf("error parsing document from document server: %v", err)
	}

	return &doc, nil
}
