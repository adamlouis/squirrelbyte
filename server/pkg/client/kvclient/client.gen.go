// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: --component client --config ../../config/api.kv.yml --package kvclient --out ./kvclient/client.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/kvmodel
package kvclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/kvmodel"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Client interface {
	ListKVs(ctx context.Context, queryParams *kvmodel.ListKVsRequest) (*kvmodel.ListKVsResponse, int, error)
	GetKV(ctx context.Context, pathParams *kvmodel.GetKVPathParams) (*kvmodel.KV, int, error)
	PutKV(ctx context.Context, pathParams *kvmodel.PutKVPathParams, body *kvmodel.KV) (*kvmodel.KV, int, error)
	DeleteKV(ctx context.Context, pathParams *kvmodel.DeleteKVPathParams) (int, error)
}

func NewHTTPClient(baseURL string) Client {
	return &client{
		baseURL: baseURL,
	}
}

type client struct {
	baseURL string
}

func (c *client) ListKVs(ctx context.Context, queryParams *kvmodel.ListKVsRequest) (*kvmodel.ListKVsResponse, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/kvs", c.baseURL))
	if err != nil {
		return nil, -1, err
	}
	u.Query().Add("page_token", queryParams.PageToken)
	u.Query().Add("page_size", strconv.Itoa(queryParams.PageSize))
	var requestBody io.Reader
	req, err := http.NewRequest(http.MethodGet, u.String(), requestBody)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		respBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, resp.StatusCode, fmt.Errorf("[%d] %s", resp.StatusCode, string(respBytes))
	}
	respBody := kvmodel.ListKVsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) GetKV(ctx context.Context, pathParams *kvmodel.GetKVPathParams) (*kvmodel.KV, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/kvs/%v", c.baseURL, pathParams.Key))
	if err != nil {
		return nil, -1, err
	}
	var requestBody io.Reader
	req, err := http.NewRequest(http.MethodGet, u.String(), requestBody)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		respBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, resp.StatusCode, fmt.Errorf("[%d] %s", resp.StatusCode, string(respBytes))
	}
	respBody := kvmodel.KV{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) PutKV(ctx context.Context, pathParams *kvmodel.PutKVPathParams, body *kvmodel.KV) (*kvmodel.KV, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/kvs/%v", c.baseURL, pathParams.Key))
	if err != nil {
		return nil, -1, err
	}
	var requestBody io.Reader
	if jsonBytes, err := json.Marshal(body); err != nil {
		return nil, -1, err
	} else {
		requestBody = bytes.NewBuffer(jsonBytes)
	}
	req, err := http.NewRequest(http.MethodPut, u.String(), requestBody)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		respBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, resp.StatusCode, fmt.Errorf("[%d] %s", resp.StatusCode, string(respBytes))
	}
	respBody := kvmodel.KV{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) DeleteKV(ctx context.Context, pathParams *kvmodel.DeleteKVPathParams) (int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/kvs/%v", c.baseURL, pathParams.Key))
	if err != nil {
		return -1, err
	}
	var requestBody io.Reader
	req, err := http.NewRequest(http.MethodDelete, u.String(), requestBody)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		respBytes, _ := ioutil.ReadAll(resp.Body)
		return resp.StatusCode, fmt.Errorf("[%d] %s", resp.StatusCode, string(respBytes))
	}
	return resp.StatusCode, nil
}
