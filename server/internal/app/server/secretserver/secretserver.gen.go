// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: '--component server --config ../../../../config/api.secret.yml --package secretserver --out-dir . --out ./secretserver.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/secretmodel'

package secretserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/secretmodel"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type HTTPHandler interface {
	ListSecrets(w http.ResponseWriter, req *http.Request)
	GetSecret(w http.ResponseWriter, req *http.Request)
	PutSecret(w http.ResponseWriter, req *http.Request)
}
type APIHandler interface {
	ListSecrets(ctx context.Context, queryParams *secretmodel.ListSecretsRequest) (*secretmodel.ListSecretsResponse, int, error)
	GetSecret(ctx context.Context, pathParams *secretmodel.GetSecretPathParams) (*secretmodel.Secret, int, error)
	PutSecret(ctx context.Context, pathParams *secretmodel.PutSecretPathParams, body *secretmodel.Secret) (*secretmodel.Secret, int, error)
}

func RegisterRouter(apiHandler APIHandler, r *mux.Router) {
	h := apiHandlerToHTTPHandler(apiHandler)
	r.Handle("/secrets", http.HandlerFunc(h.ListSecrets)).Methods(http.MethodGet)
	r.Handle("/secrets/{key}", http.HandlerFunc(h.GetSecret)).Methods(http.MethodGet)
	r.Handle("/secrets/{key}", http.HandlerFunc(h.PutSecret)).Methods(http.MethodPut)
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

func (h *httpHandler) ListSecrets(w http.ResponseWriter, req *http.Request) {
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
	queryParams := secretmodel.ListSecretsRequest{
		PageSize:  pageSizeQueryParam,
		PageToken: pageTokenQueryParam,
	}
	r, code, err := h.apiHandler.ListSecrets(req.Context(), &queryParams)
	if err != nil {
		sendError(w, code, err)
		return
	}
	sendOK(w, code, r)
}
func (h *httpHandler) GetSecret(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key, ok := vars["key"]
	if !ok {
		sendError(w, http.StatusInternalServerError, fmt.Errorf("invalid key path parameter"))
		return
	}
	pathParams := secretmodel.GetSecretPathParams{
		Key: key,
	}
	r, code, err := h.apiHandler.GetSecret(req.Context(), &pathParams)
	if err != nil {
		sendError(w, code, err)
		return
	}
	sendOK(w, code, r)
}
func (h *httpHandler) PutSecret(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key, ok := vars["key"]
	if !ok {
		sendError(w, http.StatusInternalServerError, fmt.Errorf("invalid key path parameter"))
		return
	}
	pathParams := secretmodel.PutSecretPathParams{
		Key: key,
	}
	var requestBody secretmodel.Secret
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}
	r, code, err := h.apiHandler.PutSecret(req.Context(), &pathParams, &requestBody)
	if err != nil {
		sendError(w, code, err)
		return
	}
	sendOK(w, code, r)
}
