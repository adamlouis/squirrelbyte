// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: '--component server --config ../../../../config/api.scheduler.yml --package schedulerserver --out-dir . --out ./schedulerserver.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/schedulermodel'

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
	PutScheduler(w http.ResponseWriter, req *http.Request)
	DeleteScheduler(w http.ResponseWriter, req *http.Request)
	GetScheduler(w http.ResponseWriter, req *http.Request)
}
type APIHandler interface {
	GetScheduler(ctx context.Context, pathParams *schedulermodel.GetSchedulerPathParams) (*schedulermodel.Scheduler, int, error)
	PutScheduler(ctx context.Context, pathParams *schedulermodel.PutSchedulerPathParams, body *schedulermodel.Scheduler) (*schedulermodel.Scheduler, int, error)
	DeleteScheduler(ctx context.Context, pathParams *schedulermodel.DeleteSchedulerPathParams) (int, error)
	ListSchedulers(ctx context.Context, queryParams *schedulermodel.ListSchedulersRequest) (*schedulermodel.ListSchedulersResponse, int, error)
	PostScheduler(ctx context.Context, body *schedulermodel.Scheduler) (*schedulermodel.Scheduler, int, error)
}

func RegisterRouter(apiHandler APIHandler, r *mux.Router) {
	h := apiHandlerToHTTPHandler(apiHandler)
	r.Handle("/schedulers", http.HandlerFunc(h.ListSchedulers)).Methods(http.MethodGet)
	r.Handle("/schedulers", http.HandlerFunc(h.PostScheduler)).Methods(http.MethodPost)
	r.Handle("/schedulers/{schedulerID}", http.HandlerFunc(h.GetScheduler)).Methods(http.MethodGet)
	r.Handle("/schedulers/{schedulerID}", http.HandlerFunc(h.PutScheduler)).Methods(http.MethodPut)
	r.Handle("/schedulers/{schedulerID}", http.HandlerFunc(h.DeleteScheduler)).Methods(http.MethodDelete)
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

func (h *httpHandler) ListSchedulers(w http.ResponseWriter, req *http.Request) {
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
	queryParams := schedulermodel.ListSchedulersRequest{
		PageToken: pageTokenQueryParam,
		PageSize:  pageSizeQueryParam,
	}
	r, code, err := h.apiHandler.ListSchedulers(req.Context(), &queryParams)
	if err != nil {
		sendError(w, code, err)
		return
	}
	sendOK(w, code, r)
}
func (h *httpHandler) PostScheduler(w http.ResponseWriter, req *http.Request) {
	var requestBody schedulermodel.Scheduler
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}
	r, code, err := h.apiHandler.PostScheduler(req.Context(), &requestBody)
	if err != nil {
		sendError(w, code, err)
		return
	}
	sendOK(w, code, r)
}
func (h *httpHandler) GetScheduler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	schedulerID, ok := vars["schedulerID"]
	if !ok {
		sendError(w, http.StatusInternalServerError, fmt.Errorf("invalid schedulerID path parameter"))
		return
	}
	pathParams := schedulermodel.GetSchedulerPathParams{
		SchedulerID: schedulerID,
	}
	r, code, err := h.apiHandler.GetScheduler(req.Context(), &pathParams)
	if err != nil {
		sendError(w, code, err)
		return
	}
	sendOK(w, code, r)
}
func (h *httpHandler) PutScheduler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	schedulerID, ok := vars["schedulerID"]
	if !ok {
		sendError(w, http.StatusInternalServerError, fmt.Errorf("invalid schedulerID path parameter"))
		return
	}
	pathParams := schedulermodel.PutSchedulerPathParams{
		SchedulerID: schedulerID,
	}
	var requestBody schedulermodel.Scheduler
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}
	r, code, err := h.apiHandler.PutScheduler(req.Context(), &pathParams, &requestBody)
	if err != nil {
		sendError(w, code, err)
		return
	}
	sendOK(w, code, r)
}
func (h *httpHandler) DeleteScheduler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	schedulerID, ok := vars["schedulerID"]
	if !ok {
		sendError(w, http.StatusInternalServerError, fmt.Errorf("invalid schedulerID path parameter"))
		return
	}
	pathParams := schedulermodel.DeleteSchedulerPathParams{
		SchedulerID: schedulerID,
	}
	code, err := h.apiHandler.DeleteScheduler(req.Context(), &pathParams)
	if err != nil {
		sendError(w, code, err)
		return
	}
	sendOK(w, code, struct{}{})
}
