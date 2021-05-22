// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: --component client --config ../../config/api.scheduler.yml --package schedulerclient --out ./schedulerclient/client.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/schedulermodel
package schedulerclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/schedulermodel"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Client interface {
	ListSchedulers(ctx context.Context, queryParams *schedulermodel.ListSchedulersRequest) (*schedulermodel.ListSchedulersResponse, int, error)
	PostScheduler(ctx context.Context, body *schedulermodel.Scheduler) (*schedulermodel.Scheduler, int, error)
	GetScheduler(ctx context.Context, pathParams *schedulermodel.GetSchedulerPathParams) (*schedulermodel.Scheduler, int, error)
	PutScheduler(ctx context.Context, pathParams *schedulermodel.PutSchedulerPathParams, body *schedulermodel.Scheduler) (*schedulermodel.Scheduler, int, error)
	DeleteScheduler(ctx context.Context, pathParams *schedulermodel.DeleteSchedulerPathParams) (int, error)
}

func NewHTTPClient(baseURL string) Client {
	return &client{
		baseURL: baseURL,
	}
}

type client struct {
	baseURL string
}

func (c *client) ListSchedulers(ctx context.Context, queryParams *schedulermodel.ListSchedulersRequest) (*schedulermodel.ListSchedulersResponse, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/schedulers", c.baseURL))
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
	respBody := schedulermodel.ListSchedulersResponse{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) PostScheduler(ctx context.Context, body *schedulermodel.Scheduler) (*schedulermodel.Scheduler, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/schedulers", c.baseURL))
	if err != nil {
		return nil, -1, err
	}
	var requestBody io.Reader
	if jsonBytes, err := json.Marshal(body); err != nil {
		return nil, -1, err
	} else {
		requestBody = bytes.NewBuffer(jsonBytes)
	}
	req, err := http.NewRequest(http.MethodPost, u.String(), requestBody)
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
	respBody := schedulermodel.Scheduler{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) GetScheduler(ctx context.Context, pathParams *schedulermodel.GetSchedulerPathParams) (*schedulermodel.Scheduler, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/schedulers/%v", c.baseURL, pathParams.SchedulerID))
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
	respBody := schedulermodel.Scheduler{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) PutScheduler(ctx context.Context, pathParams *schedulermodel.PutSchedulerPathParams, body *schedulermodel.Scheduler) (*schedulermodel.Scheduler, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/schedulers/%v", c.baseURL, pathParams.SchedulerID))
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
	respBody := schedulermodel.Scheduler{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) DeleteScheduler(ctx context.Context, pathParams *schedulermodel.DeleteSchedulerPathParams) (int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/schedulers/%v", c.baseURL, pathParams.SchedulerID))
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
