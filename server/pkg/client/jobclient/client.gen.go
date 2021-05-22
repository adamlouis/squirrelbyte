// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: --component client --config ../../config/api.job.yml --package jobclient --out ./jobclient/client.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel
package jobclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Client interface {
	ListJobs(ctx context.Context, queryParams *jobmodel.ListJobsQueryParams) (*jobmodel.ListJobsResponse, int, error)
	GetJob(ctx context.Context, pathParams *jobmodel.GetJobPathParams) (*jobmodel.Job, int, error)
	DeleteJob(ctx context.Context, pathParams *jobmodel.DeleteJobPathParams) (int, error)
	QueueJob(ctx context.Context, body *jobmodel.Job) (*jobmodel.Job, int, error)
	ClaimSomeJob(ctx context.Context, body *jobmodel.ClaimSomeJobRequest) (*jobmodel.Job, int, error)
	ClaimJob(ctx context.Context, pathParams *jobmodel.ClaimJobPathParams) (*jobmodel.Job, int, error)
	ReleaseJob(ctx context.Context, pathParams *jobmodel.ReleaseJobPathParams) (*jobmodel.Job, int, error)
	SetJobSuccess(ctx context.Context, pathParams *jobmodel.SetJobSuccessPathParams) (*jobmodel.Job, int, error)
	SetJobError(ctx context.Context, pathParams *jobmodel.SetJobErrorPathParams) (*jobmodel.Job, int, error)
}

func NewHTTPClient(baseURL string) Client {
	return &client{
		baseURL: baseURL,
	}
}

type client struct {
	baseURL string
}

func (c *client) ListJobs(ctx context.Context, queryParams *jobmodel.ListJobsQueryParams) (*jobmodel.ListJobsResponse, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/jobs", c.baseURL))
	if err != nil {
		return nil, -1, err
	}
	u.Query().Add("page_size", strconv.Itoa(queryParams.PageSize))
	u.Query().Add("page_token", queryParams.PageToken)
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
	respBody := jobmodel.ListJobsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) GetJob(ctx context.Context, pathParams *jobmodel.GetJobPathParams) (*jobmodel.Job, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/jobs/%v", c.baseURL, pathParams.JobID))
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
	respBody := jobmodel.Job{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) DeleteJob(ctx context.Context, pathParams *jobmodel.DeleteJobPathParams) (int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/jobs/%v", c.baseURL, pathParams.JobID))
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
func (c *client) QueueJob(ctx context.Context, body *jobmodel.Job) (*jobmodel.Job, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/jobs:queue", c.baseURL))
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
	respBody := jobmodel.Job{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) ClaimSomeJob(ctx context.Context, body *jobmodel.ClaimSomeJobRequest) (*jobmodel.Job, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/jobs:claim", c.baseURL))
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
	respBody := jobmodel.Job{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) ClaimJob(ctx context.Context, pathParams *jobmodel.ClaimJobPathParams) (*jobmodel.Job, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/jobs/%v:claim", c.baseURL, pathParams.JobID))
	if err != nil {
		return nil, -1, err
	}
	var requestBody io.Reader
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
	respBody := jobmodel.Job{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) ReleaseJob(ctx context.Context, pathParams *jobmodel.ReleaseJobPathParams) (*jobmodel.Job, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/jobs/%v:release", c.baseURL, pathParams.JobID))
	if err != nil {
		return nil, -1, err
	}
	var requestBody io.Reader
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
	respBody := jobmodel.Job{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) SetJobSuccess(ctx context.Context, pathParams *jobmodel.SetJobSuccessPathParams) (*jobmodel.Job, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/jobs/%v:success", c.baseURL, pathParams.JobID))
	if err != nil {
		return nil, -1, err
	}
	var requestBody io.Reader
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
	respBody := jobmodel.Job{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) SetJobError(ctx context.Context, pathParams *jobmodel.SetJobErrorPathParams) (*jobmodel.Job, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/jobs/%v:error", c.baseURL, pathParams.JobID))
	if err != nil {
		return nil, -1, err
	}
	var requestBody io.Reader
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
	respBody := jobmodel.Job{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
