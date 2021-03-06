// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: --component server --config ../../../../config/api.job.yml --package jobserver --out-dir . --out ./jobserver.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel
package jobserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type HTTPHandler interface {
	ListJobs(w http.ResponseWriter, req *http.Request)
	GetJob(w http.ResponseWriter, req *http.Request)
	DeleteJob(w http.ResponseWriter, req *http.Request)
	QueueJob(w http.ResponseWriter, req *http.Request)
	ClaimSomeJob(w http.ResponseWriter, req *http.Request)
	ClaimJob(w http.ResponseWriter, req *http.Request)
	ReleaseJob(w http.ResponseWriter, req *http.Request)
	SetJobSuccess(w http.ResponseWriter, req *http.Request)
	SetJobError(w http.ResponseWriter, req *http.Request)
}
type APIHandler interface {
	ListJobs(ctx context.Context, queryParams *jobmodel.ListJobsQueryParams) (*jobmodel.ListJobsResponse, error)
	GetJob(ctx context.Context, pathParams *jobmodel.GetJobPathParams) (*jobmodel.Job, error)
	DeleteJob(ctx context.Context, pathParams *jobmodel.DeleteJobPathParams) error
	QueueJob(ctx context.Context, body *jobmodel.Job) (*jobmodel.Job, error)
	ClaimSomeJob(ctx context.Context, body *jobmodel.ClaimSomeJobRequest) (*jobmodel.Job, error)
	ClaimJob(ctx context.Context, pathParams *jobmodel.ClaimJobPathParams) (*jobmodel.Job, error)
	ReleaseJob(ctx context.Context, pathParams *jobmodel.ReleaseJobPathParams) (*jobmodel.Job, error)
	SetJobSuccess(ctx context.Context, pathParams *jobmodel.SetJobSuccessPathParams) (*jobmodel.Job, error)
	SetJobError(ctx context.Context, pathParams *jobmodel.SetJobErrorPathParams) (*jobmodel.Job, error)
}

func RegisterRouter(apiHandler APIHandler, r *mux.Router, c ErrorCoder) {
	h := apiHandlerToHTTPHandler(apiHandler, c)
	r.Handle("/jobs", http.HandlerFunc(h.ListJobs)).Methods(http.MethodGet)
	r.Handle("/jobs/{jobID}", http.HandlerFunc(h.GetJob)).Methods(http.MethodGet)
	r.Handle("/jobs/{jobID}", http.HandlerFunc(h.DeleteJob)).Methods(http.MethodDelete)
	r.Handle("/jobs:queue", http.HandlerFunc(h.QueueJob)).Methods(http.MethodPost)
	r.Handle("/jobs:claim", http.HandlerFunc(h.ClaimSomeJob)).Methods(http.MethodPost)
	r.Handle("/jobs/{jobID}:claim", http.HandlerFunc(h.ClaimJob)).Methods(http.MethodPost)
	r.Handle("/jobs/{jobID}:release", http.HandlerFunc(h.ReleaseJob)).Methods(http.MethodPost)
	r.Handle("/jobs/{jobID}:success", http.HandlerFunc(h.SetJobSuccess)).Methods(http.MethodPost)
	r.Handle("/jobs/{jobID}:error", http.HandlerFunc(h.SetJobError)).Methods(http.MethodPost)
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

func (h *httpHandler) ListJobs(w http.ResponseWriter, req *http.Request) {
	pageSizeQueryParam := 0
	if req.URL.Query().Get("page_size") != "" {
		q, err := strconv.Atoi(req.URL.Query().Get("page_size"))
		if err != nil {
			sendErrorWithCode(w, http.StatusBadRequest, err)
			return
		}
		pageSizeQueryParam = q
	}
	pageTokenQueryParam := req.URL.Query().Get("page_token")
	queryParams := jobmodel.ListJobsQueryParams{
		PageSize:  pageSizeQueryParam,
		PageToken: pageTokenQueryParam,
	}
	r, err := h.apiHandler.ListJobs(req.Context(), &queryParams)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) GetJob(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jobID, ok := vars["jobID"]
	if !ok {
		sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid jobID path parameter"))
		return
	}
	pathParams := jobmodel.GetJobPathParams{
		JobID: jobID,
	}
	r, err := h.apiHandler.GetJob(req.Context(), &pathParams)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) DeleteJob(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jobID, ok := vars["jobID"]
	if !ok {
		sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid jobID path parameter"))
		return
	}
	pathParams := jobmodel.DeleteJobPathParams{
		JobID: jobID,
	}
	err := h.apiHandler.DeleteJob(req.Context(), &pathParams)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, struct{}{})
}
func (h *httpHandler) QueueJob(w http.ResponseWriter, req *http.Request) {
	var requestBody jobmodel.Job
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		sendErrorWithCode(w, http.StatusBadRequest, err)
		return
	}
	r, err := h.apiHandler.QueueJob(req.Context(), &requestBody)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) ClaimSomeJob(w http.ResponseWriter, req *http.Request) {
	var requestBody jobmodel.ClaimSomeJobRequest
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		sendErrorWithCode(w, http.StatusBadRequest, err)
		return
	}
	r, err := h.apiHandler.ClaimSomeJob(req.Context(), &requestBody)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) ClaimJob(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jobID, ok := vars["jobID"]
	if !ok {
		sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid jobID path parameter"))
		return
	}
	pathParams := jobmodel.ClaimJobPathParams{
		JobID: jobID,
	}
	r, err := h.apiHandler.ClaimJob(req.Context(), &pathParams)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) ReleaseJob(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jobID, ok := vars["jobID"]
	if !ok {
		sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid jobID path parameter"))
		return
	}
	pathParams := jobmodel.ReleaseJobPathParams{
		JobID: jobID,
	}
	r, err := h.apiHandler.ReleaseJob(req.Context(), &pathParams)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) SetJobSuccess(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jobID, ok := vars["jobID"]
	if !ok {
		sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid jobID path parameter"))
		return
	}
	pathParams := jobmodel.SetJobSuccessPathParams{
		JobID: jobID,
	}
	r, err := h.apiHandler.SetJobSuccess(req.Context(), &pathParams)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) SetJobError(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jobID, ok := vars["jobID"]
	if !ok {
		sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid jobID path parameter"))
		return
	}
	pathParams := jobmodel.SetJobErrorPathParams{
		JobID: jobID,
	}
	r, err := h.apiHandler.SetJobError(req.Context(), &pathParams)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
