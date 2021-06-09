// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: --component client --config ../../config/api.oauth.yml --package oauthclient --out ./oauthclient/client.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/oauthmodel
package oauthclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/oauthmodel"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Client interface {
	ListProviders(ctx context.Context, queryParams *oauthmodel.ListOAuthProvidersRequest) (*oauthmodel.ListOAuthProvidersResponse, int, error)
	GetOAuthAuthorizationURL(ctx context.Context, pathParams *oauthmodel.GetOAuthAuthorizationURLPathParams) (*oauthmodel.GetOAuthAuthorizationURLResponse, int, error)
	GetOAuthToken(ctx context.Context, pathParams *oauthmodel.GetOAuthTokenPathParams, body *oauthmodel.GetOAuthTokenRequest) (*oauthmodel.Token, int, error)
	GetOAuthConfig(ctx context.Context, pathParams *oauthmodel.GetOAuthConfigPathParams) (*oauthmodel.Config, int, error)
	PutOAuthConfig(ctx context.Context, pathParams *oauthmodel.PutOAuthConfigPathParams, body *oauthmodel.Config) (*oauthmodel.Config, int, error)
	DeleteOAuthConfig(ctx context.Context, pathParams *oauthmodel.DeleteOAuthConfigPathParams) (int, error)
}

func NewHTTPClient(baseURL string) Client {
	return &client{
		baseURL: baseURL,
	}
}

type client struct {
	baseURL string
}

func (c *client) ListProviders(ctx context.Context, queryParams *oauthmodel.ListOAuthProvidersRequest) (*oauthmodel.ListOAuthProvidersResponse, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/oauth/providers", c.baseURL))
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
	respBody := oauthmodel.ListOAuthProvidersResponse{}
	if len(respBytes) == 0 {
		return nil, resp.StatusCode, nil
	}
	err = json.Unmarshal(respBytes, &respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) GetOAuthAuthorizationURL(ctx context.Context, pathParams *oauthmodel.GetOAuthAuthorizationURLPathParams) (*oauthmodel.GetOAuthAuthorizationURLResponse, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/oauth/providers/%v/authorize", c.baseURL, pathParams.Name))
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
	respBody := oauthmodel.GetOAuthAuthorizationURLResponse{}
	if len(respBytes) == 0 {
		return nil, resp.StatusCode, nil
	}
	err = json.Unmarshal(respBytes, &respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) GetOAuthToken(ctx context.Context, pathParams *oauthmodel.GetOAuthTokenPathParams, body *oauthmodel.GetOAuthTokenRequest) (*oauthmodel.Token, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/oauth/providers/%v/token", c.baseURL, pathParams.Name))
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
	respBody := oauthmodel.Token{}
	if len(respBytes) == 0 {
		return nil, resp.StatusCode, nil
	}
	err = json.Unmarshal(respBytes, &respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) GetOAuthConfig(ctx context.Context, pathParams *oauthmodel.GetOAuthConfigPathParams) (*oauthmodel.Config, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/oauth/providers/%v/config", c.baseURL, pathParams.Name))
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
	respBody := oauthmodel.Config{}
	if len(respBytes) == 0 {
		return nil, resp.StatusCode, nil
	}
	err = json.Unmarshal(respBytes, &respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) PutOAuthConfig(ctx context.Context, pathParams *oauthmodel.PutOAuthConfigPathParams, body *oauthmodel.Config) (*oauthmodel.Config, int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/oauth/providers/%v/config", c.baseURL, pathParams.Name))
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
	respBody := oauthmodel.Config{}
	if len(respBytes) == 0 {
		return nil, resp.StatusCode, nil
	}
	err = json.Unmarshal(respBytes, &respBody)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return &respBody, resp.StatusCode, nil
}
func (c *client) DeleteOAuthConfig(ctx context.Context, pathParams *oauthmodel.DeleteOAuthConfigPathParams) (int, error) {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/oauth/providers/%v/config", c.baseURL, pathParams.Name))
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
