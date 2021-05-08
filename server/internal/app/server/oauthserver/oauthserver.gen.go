// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: '--component server --config ../../../../config/api.oauth.yml --package oauthserver --out-dir . --out ./oauthserver.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/oauthmodel'

package oauthserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/oauthmodel"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type HTTPHandler interface {
	ListProviders(w http.ResponseWriter, req *http.Request)
	GetOAuthAuthorizationURL(w http.ResponseWriter, req *http.Request)
	GetOAuthToken(w http.ResponseWriter, req *http.Request)
	GetOAuthConfig(w http.ResponseWriter, req *http.Request)
	PutOAuthConfig(w http.ResponseWriter, req *http.Request)
}
type APIHandler interface {
	GetOAuthConfig(ctx context.Context, pathParams *oauthmodel.GetOAuthConfigPathParams, body *oauthmodel.Config) (*oauthmodel.Config, int, error)
	PutOAuthConfig(ctx context.Context, pathParams *oauthmodel.PutOAuthConfigPathParams, body *oauthmodel.Config) (*oauthmodel.Config, int, error)
	ListProviders(ctx context.Context, queryParams *oauthmodel.ListOAuthProvidersRequest) (*oauthmodel.ListOAuthProvidersResponse, int, error)
	GetOAuthAuthorizationURL(ctx context.Context, pathParams *oauthmodel.GetOAuthAuthorizationURLPathParams) (*oauthmodel.GetOAuthAuthorizationURLResponse, int, error)
	GetOAuthToken(ctx context.Context, pathParams *oauthmodel.GetOAuthTokenPathParams, body *oauthmodel.GetOAuthTokenRequest) (*oauthmodel.Token, int, error)
}

func RegisterRouter(apiHandler APIHandler, r *mux.Router) {
	h := apiHandlerToHTTPHandler(apiHandler)
	r.Handle("/oauth/providers", http.HandlerFunc(h.ListProviders)).Methods(http.MethodGet)
	r.Handle("/oauth/providers/{name}/authorize", http.HandlerFunc(h.GetOAuthAuthorizationURL)).Methods(http.MethodGet)
	r.Handle("/oauth/providers/{name}/token", http.HandlerFunc(h.GetOAuthToken)).Methods(http.MethodPost)
	r.Handle("/oauth/providers/{name}/config", http.HandlerFunc(h.GetOAuthConfig)).Methods(http.MethodGet)
	r.Handle("/oauth/providers/{name}/config", http.HandlerFunc(h.PutOAuthConfig)).Methods(http.MethodPut)
}
func apiHandlerToHTTPHandler(apiHandler APIHandler) HTTPHandler {
	return &httpHandler{
		apiHandler: apiHandler,
	}
}

type httpHandler struct {
	apiHandler APIHandler
}

// sendError sends an error response
func sendError(w http.ResponseWriter, code int, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	e := json.NewEncoder(w)
	e.SetEscapeHTML(false)
	e.Encode(&errorResponse{
		Message: err.Error(),
	})
}

// sendOK sends an success response
func sendOK(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	e := json.NewEncoder(w)
	e.SetEscapeHTML(false)
	e.Encode(body)
}

type errorResponse struct {
	Message string `json:"message"`
}

func (h *httpHandler) GetOAuthConfig(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		sendError(w, http.StatusInternalServerError, fmt.Errorf("invalid name path parameter"))
		return
	}
	pathParams := oauthmodel.GetOAuthConfigPathParams{
		Name: name,
	}
	var requestBody oauthmodel.Config
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}
	r, code, err := h.apiHandler.GetOAuthConfig(req.Context(), &pathParams, &requestBody)
	if err != nil {
		sendError(w, code, err)
		return
	}
	sendOK(w, code, r)
}
func (h *httpHandler) PutOAuthConfig(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		sendError(w, http.StatusInternalServerError, fmt.Errorf("invalid name path parameter"))
		return
	}
	pathParams := oauthmodel.PutOAuthConfigPathParams{
		Name: name,
	}
	var requestBody oauthmodel.Config
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}
	r, code, err := h.apiHandler.PutOAuthConfig(req.Context(), &pathParams, &requestBody)
	if err != nil {
		sendError(w, code, err)
		return
	}
	sendOK(w, code, r)
}
func (h *httpHandler) ListProviders(w http.ResponseWriter, req *http.Request) {
	pageTokenQueryParam := req.URL.Query().Get("page_token")
	pageSizeQueryParam := 0
	if req.URL.Query().Get("page_size") != "" {
		q, err := strconv.Atoi(req.URL.Query().Get("page_size"))
		if err != nil {
			sendError(w, http.StatusBadRequest, err)
			return
		}
		pageSizeQueryParam = q
	}
	queryParams := oauthmodel.ListOAuthProvidersRequest{
		PageToken: pageTokenQueryParam,
		PageSize:  pageSizeQueryParam,
	}
	r, code, err := h.apiHandler.ListProviders(req.Context(), &queryParams)
	if err != nil {
		sendError(w, code, err)
		return
	}
	sendOK(w, code, r)
}
func (h *httpHandler) GetOAuthAuthorizationURL(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		sendError(w, http.StatusInternalServerError, fmt.Errorf("invalid name path parameter"))
		return
	}
	pathParams := oauthmodel.GetOAuthAuthorizationURLPathParams{
		Name: name,
	}
	r, code, err := h.apiHandler.GetOAuthAuthorizationURL(req.Context(), &pathParams)
	if err != nil {
		sendError(w, code, err)
		return
	}
	sendOK(w, code, r)
}
func (h *httpHandler) GetOAuthToken(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		sendError(w, http.StatusInternalServerError, fmt.Errorf("invalid name path parameter"))
		return
	}
	pathParams := oauthmodel.GetOAuthTokenPathParams{
		Name: name,
	}
	var requestBody oauthmodel.GetOAuthTokenRequest
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}
	r, code, err := h.apiHandler.GetOAuthToken(req.Context(), &pathParams, &requestBody)
	if err != nil {
		sendError(w, code, err)
		return
	}
	sendOK(w, code, r)
}
