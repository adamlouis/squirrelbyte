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

type JobClient interface {
	Queue(ctx context.Context, name string, input map[string]interface{}) error
	Claim(ctx context.Context, opts *model.ClaimJobRequest) (*model.Job, error)
	Release(ctx context.Context, id string) error
	SetSuccess(ctx context.Context, id string, output map[string]interface{}) error
	SetError(ctx context.Context, id string, output map[string]interface{}) error
}

func NewHTTPJobClient(url string) JobClient {
	return &jobClient{
		url: url,
	}
}

type jobClient struct {
	url string
}

func (jc *jobClient) Queue(ctx context.Context, name string, input map[string]interface{}) error {
	b, err := json.Marshal(model.Job{
		Name:  name,
		Input: input,
	})
	if err != nil {
		return err
	}

	res, err := http.Post(fmt.Sprintf("%s/api/jobs:queue", jc.url), "application/json", bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("error setting success: %v", err)
	}

	if res.StatusCode != 200 {
		b, _ := ioutil.ReadAll(res.Body)
		return fmt.Errorf("error setting success: %s: %s", res.Status, string(b))
	}

	return nil
}

func (jc *jobClient) SetSuccess(ctx context.Context, id string, output map[string]interface{}) error {
	b, err := json.Marshal(model.Job{
		Output: output,
	})
	if err != nil {
		return err
	}

	res, err := http.Post(fmt.Sprintf("%s/api/jobs/%s:success", jc.url, id), "application/json", bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("error setting success: %v", err)
	}

	if res.StatusCode != 200 {
		b, _ := ioutil.ReadAll(res.Body)
		return fmt.Errorf("error setting success: %s: %s", res.Status, string(b))
	}

	return nil
}

func (jc *jobClient) SetError(ctx context.Context, id string, output map[string]interface{}) error {
	b, err := json.Marshal(model.Job{
		Output: output,
	})
	if err != nil {
		return err
	}

	res, err := http.Post(
		fmt.Sprintf("%s/api/jobs/%s:error", jc.url, id),
		"application/json",
		bytes.NewBuffer(b),
	)
	if err != nil {
		return fmt.Errorf("error setting error: %v", err)
	}

	if res.StatusCode != 200 {
		b, _ := ioutil.ReadAll(res.Body)
		return fmt.Errorf("error setting error: %s: %s", res.Status, string(b))
	}

	return nil
}

func (jc *jobClient) Release(ctx context.Context, id string) error {
	res, err := http.Post(
		fmt.Sprintf("%s/api/jobs/%s:release", jc.url, id),
		"application/json",
		bytes.NewBuffer([]byte("{}")))
	if err != nil {
		return fmt.Errorf("error releasing: %v", err)
	}

	if res.StatusCode != 200 {
		b, _ := ioutil.ReadAll(res.Body)
		return fmt.Errorf("error releasing: %s: %s", res.Status, string(b))
	}

	return nil
}
func (jc *jobClient) Claim(ctx context.Context, opts *model.ClaimJobRequest) (*model.Job, error) {
	b, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(
		fmt.Sprintf("%s/api/jobs:claim", jc.url),
		"application/json",
		bytes.NewBuffer(b),
	)
	if err != nil {
		return nil, fmt.Errorf("error claiming job from job server: %v", err)
	}

	if res.StatusCode != 200 {
		b, _ := ioutil.ReadAll(res.Body)
		return nil, fmt.Errorf("error claiming job from job server: %s: %s", res.Status, string(b))
	}

	j := model.Job{}
	err = json.NewDecoder(res.Body).Decode(&j)
	if err != nil {
		return nil, fmt.Errorf("error parsing job from job server: %v", err)
	}

	return &j, nil
}
