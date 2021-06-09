// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: --component server --config ../../../../config/api.kv.yml --package kvserver --out-dir . --out ./kvserver.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/kvmodel
package kvserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/kvmodel"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type HTTPHandler interface {
	ListKVs(w http.ResponseWriter, req *http.Request)
	GetKV(w http.ResponseWriter, req *http.Request)
	PutKV(w http.ResponseWriter, req *http.Request)
	DeleteKV(w http.ResponseWriter, req *http.Request)
}
type APIHandler interface {
	ListKVs(ctx context.Context, queryParams *kvmodel.ListKVsRequest) (*kvmodel.ListKVsResponse, error)
	GetKV(ctx context.Context, pathParams *kvmodel.GetKVPathParams) (*kvmodel.KV, error)
	PutKV(ctx context.Context, pathParams *kvmodel.PutKVPathParams, body *kvmodel.KV) (*kvmodel.KV, error)
	DeleteKV(ctx context.Context, pathParams *kvmodel.DeleteKVPathParams) error
}

func RegisterRouter(apiHandler APIHandler, r *mux.Router, c ErrorCoder) {
	h := apiHandlerToHTTPHandler(apiHandler, c)
	r.Handle("/kvs", http.HandlerFunc(h.ListKVs)).Methods(http.MethodGet)
	r.Handle("/kvs/{key}", http.HandlerFunc(h.GetKV)).Methods(http.MethodGet)
	r.Handle("/kvs/{key}", http.HandlerFunc(h.PutKV)).Methods(http.MethodPut)
	r.Handle("/kvs/{key}", http.HandlerFunc(h.DeleteKV)).Methods(http.MethodDelete)
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

func (h *httpHandler) ListKVs(w http.ResponseWriter, req *http.Request) {
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
	queryParams := kvmodel.ListKVsRequest{
		PageToken: pageTokenQueryParam,
		PageSize:  pageSizeQueryParam,
	}
	r, err := h.apiHandler.ListKVs(req.Context(), &queryParams)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) GetKV(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key, ok := vars["key"]
	if !ok {
		sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid key path parameter"))
		return
	}
	pathParams := kvmodel.GetKVPathParams{
		Key: key,
	}
	r, err := h.apiHandler.GetKV(req.Context(), &pathParams)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) PutKV(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key, ok := vars["key"]
	if !ok {
		sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid key path parameter"))
		return
	}
	pathParams := kvmodel.PutKVPathParams{
		Key: key,
	}
	var requestBody kvmodel.KV
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		sendErrorWithCode(w, http.StatusBadRequest, err)
		return
	}
	r, err := h.apiHandler.PutKV(req.Context(), &pathParams, &requestBody)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) DeleteKV(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key, ok := vars["key"]
	if !ok {
		sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid key path parameter"))
		return
	}
	pathParams := kvmodel.DeleteKVPathParams{
		Key: key,
	}
	err := h.apiHandler.DeleteKV(req.Context(), &pathParams)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, struct{}{})
}
