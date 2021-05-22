// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: --component server --config ../../../../config/api.scheduler.yml --package schedulerserver --out-dir . --out ./schedulerserver.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/schedulermodel
package schedulerserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/schedulermodel"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type HTTPHandler interface {
	ListSchedulers(w http.ResponseWriter, req *http.Request)
	PostScheduler(w http.ResponseWriter, req *http.Request)
	GetScheduler(w http.ResponseWriter, req *http.Request)
	PutScheduler(w http.ResponseWriter, req *http.Request)
	DeleteScheduler(w http.ResponseWriter, req *http.Request)
}
type APIHandler interface {
	ListSchedulers(ctx context.Context, queryParams *schedulermodel.ListSchedulersRequest) (*schedulermodel.ListSchedulersResponse, error)
	PostScheduler(ctx context.Context, body *schedulermodel.Scheduler) (*schedulermodel.Scheduler, error)
	GetScheduler(ctx context.Context, pathParams *schedulermodel.GetSchedulerPathParams) (*schedulermodel.Scheduler, error)
	PutScheduler(ctx context.Context, pathParams *schedulermodel.PutSchedulerPathParams, body *schedulermodel.Scheduler) (*schedulermodel.Scheduler, error)
	DeleteScheduler(ctx context.Context, pathParams *schedulermodel.DeleteSchedulerPathParams) error
}

func RegisterRouter(apiHandler APIHandler, r *mux.Router, c ErrorCoder) {
	h := apiHandlerToHTTPHandler(apiHandler, c)
	r.Handle("/schedulers", http.HandlerFunc(h.ListSchedulers)).Methods(http.MethodGet)
	r.Handle("/schedulers", http.HandlerFunc(h.PostScheduler)).Methods(http.MethodPost)
	r.Handle("/schedulers/{schedulerID}", http.HandlerFunc(h.GetScheduler)).Methods(http.MethodGet)
	r.Handle("/schedulers/{schedulerID}", http.HandlerFunc(h.PutScheduler)).Methods(http.MethodPut)
	r.Handle("/schedulers/{schedulerID}", http.HandlerFunc(h.DeleteScheduler)).Methods(http.MethodDelete)
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
	e := json.NewEncoder(w)
	e.SetEscapeHTML(false)
	e.Encode(body)
}

type errorResponse struct {
	Message string `json:"message"`
}

func (h *httpHandler) ListSchedulers(w http.ResponseWriter, req *http.Request) {
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
	queryParams := schedulermodel.ListSchedulersRequest{
		PageToken: pageTokenQueryParam,
		PageSize:  pageSizeQueryParam,
	}
	r, err := h.apiHandler.ListSchedulers(req.Context(), &queryParams)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) PostScheduler(w http.ResponseWriter, req *http.Request) {
	var requestBody schedulermodel.Scheduler
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		sendErrorWithCode(w, http.StatusBadRequest, err)
		return
	}
	r, err := h.apiHandler.PostScheduler(req.Context(), &requestBody)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) GetScheduler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	schedulerID, ok := vars["schedulerID"]
	if !ok {
		sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid schedulerID path parameter"))
		return
	}
	pathParams := schedulermodel.GetSchedulerPathParams{
		SchedulerID: schedulerID,
	}
	r, err := h.apiHandler.GetScheduler(req.Context(), &pathParams)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) PutScheduler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	schedulerID, ok := vars["schedulerID"]
	if !ok {
		sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid schedulerID path parameter"))
		return
	}
	pathParams := schedulermodel.PutSchedulerPathParams{
		SchedulerID: schedulerID,
	}
	var requestBody schedulermodel.Scheduler
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		sendErrorWithCode(w, http.StatusBadRequest, err)
		return
	}
	r, err := h.apiHandler.PutScheduler(req.Context(), &pathParams, &requestBody)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) DeleteScheduler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	schedulerID, ok := vars["schedulerID"]
	if !ok {
		sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid schedulerID path parameter"))
		return
	}
	pathParams := schedulermodel.DeleteSchedulerPathParams{
		SchedulerID: schedulerID,
	}
	err := h.apiHandler.DeleteScheduler(req.Context(), &pathParams)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, struct{}{})
}
