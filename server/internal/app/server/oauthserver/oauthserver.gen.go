// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: --component server --config ../../../../config/api.oauth.yml --package oauthserver --out-dir . --out ./oauthserver.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/oauthmodel
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
	DeleteOAuthConfig(w http.ResponseWriter, req *http.Request)
}
type APIHandler interface {
	ListProviders(ctx context.Context, queryParams *oauthmodel.ListOAuthProvidersRequest) (*oauthmodel.ListOAuthProvidersResponse, error)
	GetOAuthAuthorizationURL(ctx context.Context, pathParams *oauthmodel.GetOAuthAuthorizationURLPathParams) (*oauthmodel.GetOAuthAuthorizationURLResponse, error)
	GetOAuthToken(ctx context.Context, pathParams *oauthmodel.GetOAuthTokenPathParams, body *oauthmodel.GetOAuthTokenRequest) (*oauthmodel.Token, error)
	GetOAuthConfig(ctx context.Context, pathParams *oauthmodel.GetOAuthConfigPathParams) (*oauthmodel.Config, error)
	PutOAuthConfig(ctx context.Context, pathParams *oauthmodel.PutOAuthConfigPathParams, body *oauthmodel.Config) (*oauthmodel.Config, error)
	DeleteOAuthConfig(ctx context.Context, pathParams *oauthmodel.DeleteOAuthConfigPathParams) error
}

func RegisterRouter(apiHandler APIHandler, r *mux.Router, c ErrorCoder) {
	h := apiHandlerToHTTPHandler(apiHandler, c)
	r.Handle("/oauth/providers", http.HandlerFunc(h.ListProviders)).Methods(http.MethodGet)
	r.Handle("/oauth/providers/{name}/authorize", http.HandlerFunc(h.GetOAuthAuthorizationURL)).Methods(http.MethodGet)
	r.Handle("/oauth/providers/{name}/token", http.HandlerFunc(h.GetOAuthToken)).Methods(http.MethodPost)
	r.Handle("/oauth/providers/{name}/config", http.HandlerFunc(h.GetOAuthConfig)).Methods(http.MethodGet)
	r.Handle("/oauth/providers/{name}/config", http.HandlerFunc(h.PutOAuthConfig)).Methods(http.MethodPut)
	r.Handle("/oauth/providers/{name}/config", http.HandlerFunc(h.DeleteOAuthConfig)).Methods(http.MethodDelete)
}

func apiHandlerToHTTPHandler(apiHandler APIHandler, errorCoder ErrorCoder) HTTPHandler {
	return &httpHandler{
		apiHandler: apiHandler,
		errorCoder: errorCoder,
	}
}

type httpHandler struct {
	apiHandler APIHandler
	errorCoder ErrorCoder
}

type ErrorCoder func(e error) int

// sendError sends an error response
func (h *httpHandler) sendError(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(h.errorCoder(err))
	e := json.NewEncoder(w)
	e.SetEscapeHTML(false)
	e.Encode(&errorResponse{
		Message: err.Error(),
	})
}

func sendErrorWithCode(w http.ResponseWriter, code int, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	e := json.NewEncoder(w)
	e.SetEscapeHTML(false)
	e.Encode(&errorResponse{
		Message: err.Error(),
	})
}

// sendOK sends an success response
func sendOK(w http.ResponseWriter, body interface{}) {
	w.Header().Add("Content-Type", "application/json")
	code := http.StatusOK
	if body == nil {
		code = http.StatusNoContent
	}
	w.WriteHeader(code)
	if body != nil {
		e := json.NewEncoder(w)
		e.SetEscapeHTML(false)
		e.Encode(body)
	}
}

type errorResponse struct {
	Message string `json:"message"`
}

func (h *httpHandler) ListProviders(w http.ResponseWriter, req *http.Request) {
	pageTokenQueryParam := req.URL.Query().Get("page_token")
	pageSizeQueryParam := 0
	if req.URL.Query().Get("page_size") != "" {
		q, err := strconv.Atoi(req.URL.Query().Get("page_size"))
		if err != nil {
			sendErrorWithCode(w, http.StatusBadRequest, err)
			return
		}
		pageSizeQueryParam = q
	}
	queryParams := oauthmodel.ListOAuthProvidersRequest{
		PageToken: pageTokenQueryParam,
		PageSize:  pageSizeQueryParam,
	}
	r, err := h.apiHandler.ListProviders(req.Context(), &queryParams)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) GetOAuthAuthorizationURL(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid name path parameter"))
		return
	}
	pathParams := oauthmodel.GetOAuthAuthorizationURLPathParams{
		Name: name,
	}
	r, err := h.apiHandler.GetOAuthAuthorizationURL(req.Context(), &pathParams)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) GetOAuthToken(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid name path parameter"))
		return
	}
	pathParams := oauthmodel.GetOAuthTokenPathParams{
		Name: name,
	}
	var requestBody oauthmodel.GetOAuthTokenRequest
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		sendErrorWithCode(w, http.StatusBadRequest, err)
		return
	}
	r, err := h.apiHandler.GetOAuthToken(req.Context(), &pathParams, &requestBody)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) GetOAuthConfig(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid name path parameter"))
		return
	}
	pathParams := oauthmodel.GetOAuthConfigPathParams{
		Name: name,
	}
	r, err := h.apiHandler.GetOAuthConfig(req.Context(), &pathParams)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) PutOAuthConfig(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid name path parameter"))
		return
	}
	pathParams := oauthmodel.PutOAuthConfigPathParams{
		Name: name,
	}
	var requestBody oauthmodel.Config
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		sendErrorWithCode(w, http.StatusBadRequest, err)
		return
	}
	r, err := h.apiHandler.PutOAuthConfig(req.Context(), &pathParams, &requestBody)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) DeleteOAuthConfig(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid name path parameter"))
		return
	}
	pathParams := oauthmodel.DeleteOAuthConfigPathParams{
		Name: name,
	}
	err := h.apiHandler.DeleteOAuthConfig(req.Context(), &pathParams)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, struct{}{})
}
