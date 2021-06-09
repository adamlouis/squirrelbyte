// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: --component client --config ../../config/api.document.yml --package documentclient --out ./documentclient/client.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel
package documentclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Client interface {
	ListDocuments(ctx context.Context, queryParams *documentmodel.ListDocumentsQueryParams) (*documentmodel.ListDocumentsResponse, int, error)
	PostDocument(ctx context.Context, body *documentmodel.Document) (*documentmodel.Document, int, error)
	GetDocument(ctx context.Context, pathParams *documentmodel.GetDocumentPathParams) (*documentmodel.Document, int, error)
	PutDocument(ctx context.Context, pathParams *documentmodel.PutDocumentPathParams, body *documentmodel.Document) (*documentmodel.Document, int, error)
	DeleteDocument(ctx context.Context, pathParams *documentmodel.DeleteDocumentPathParams) (int, error)
	QueryDocuments(ctx context.Context, body *documentmodel.QueryDocumentsRequest) (*documentmodel.QueryDocumentsResponse, int, error)
}

func NewHTTPClient(baseURL string) Client {
	return &client{
		baseURL: baseURL,
	}
}

type client struct {
	baseURL string
}

func (c *client) ListDocuments(ctx context.Context, queryParams *documentmodel.ListDocumentsQueryParams) (*documentmodel.ListDocumentsResponse, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/documents", c.baseURL))
	if err != nil {
		return nil, -1, err
	}
	u.Query().Add("page_token", queryParams.PageToken)
	u.Query().Add("page_size", strconv.Itoa(queryParams.PageSize))
	var requestBody io.Reader
	req, err := http.NewRequest(" + golangMethodByMethod[route.Method] + ", u.String(), requestBody)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, fmt.Errorf("[%d] %s", resp.StatusCode, string(respBytes))
	}
	respBody := documentmodel.ListDocumentsResponse{}
	if len(respBytes) == 0 {
		return nil, resp.StatusCode, nil
	}
	err = json.Unmarshal(respBytes, &respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) PostDocument(ctx context.Context, body *documentmodel.Document) (*documentmodel.Document, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/documents", c.baseURL))
	if err != nil {
		return nil, -1, err
	}
	var requestBody io.Reader
	if jsonBytes, err := json.Marshal(body); err != nil {
		return nil, -1, err
	} else {
		requestBody = bytes.NewBuffer(jsonBytes)
	}
	req, err := http.NewRequest(" + golangMethodByMethod[route.Method] + ", u.String(), requestBody)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, fmt.Errorf("[%d] %s", resp.StatusCode, string(respBytes))
	}
	respBody := documentmodel.Document{}
	if len(respBytes) == 0 {
		return nil, resp.StatusCode, nil
	}
	err = json.Unmarshal(respBytes, &respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) GetDocument(ctx context.Context, pathParams *documentmodel.GetDocumentPathParams) (*documentmodel.Document, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/documents/%v", c.baseURL, pathParams.DocumentID))
	if err != nil {
		return nil, -1, err
	}
	var requestBody io.Reader
	req, err := http.NewRequest(" + golangMethodByMethod[route.Method] + ", u.String(), requestBody)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, fmt.Errorf("[%d] %s", resp.StatusCode, string(respBytes))
	}
	respBody := documentmodel.Document{}
	if len(respBytes) == 0 {
		return nil, resp.StatusCode, nil
	}
	err = json.Unmarshal(respBytes, &respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) PutDocument(ctx context.Context, pathParams *documentmodel.PutDocumentPathParams, body *documentmodel.Document) (*documentmodel.Document, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/documents/%v", c.baseURL, pathParams.DocumentID))
	if err != nil {
		return nil, -1, err
	}
	var requestBody io.Reader
	if jsonBytes, err := json.Marshal(body); err != nil {
		return nil, -1, err
	} else {
		requestBody = bytes.NewBuffer(jsonBytes)
	}
	req, err := http.NewRequest(" + golangMethodByMethod[route.Method] + ", u.String(), requestBody)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, fmt.Errorf("[%d] %s", resp.StatusCode, string(respBytes))
	}
	respBody := documentmodel.Document{}
	if len(respBytes) == 0 {
		return nil, resp.StatusCode, nil
	}
	err = json.Unmarshal(respBytes, &respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) DeleteDocument(ctx context.Context, pathParams *documentmodel.DeleteDocumentPathParams) (int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/documents/%v", c.baseURL, pathParams.DocumentID))
	if err != nil {
		return -1, err
	}
	var requestBody io.Reader
	req, err := http.NewRequest(" + golangMethodByMethod[route.Method] + ", u.String(), requestBody)
	if err != nil {
		return -1, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}
	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, fmt.Errorf("[%d] %s", resp.StatusCode, string(respBytes))
	}
	return resp.StatusCode, nil
}
func (c *client) QueryDocuments(ctx context.Context, body *documentmodel.QueryDocumentsRequest) (*documentmodel.QueryDocumentsResponse, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/documents:query", c.baseURL))
	if err != nil {
		return nil, -1, err
	}
	var requestBody io.Reader
	if jsonBytes, err := json.Marshal(body); err != nil {
		return nil, -1, err
	} else {
		requestBody = bytes.NewBuffer(jsonBytes)
	}
	req, err := http.NewRequest(" + golangMethodByMethod[route.Method] + ", u.String(), requestBody)
	if err != nil {
		return nil, -1, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, fmt.Errorf("[%d] %s", resp.StatusCode, string(respBytes))
	}
	respBody := documentmodel.QueryDocumentsResponse{}
	if len(respBytes) == 0 {
		return nil, resp.StatusCode, nil
	}
	err = json.Unmarshal(respBytes, &respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
