package serverdef

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/adamlouis/squirrelbyte/server/pkg/model"
	"github.com/gorilla/mux"
)

// TODO: generate all of `serverdef` package from conf / openapi declaration

// convert a APIHandler to a HTTPAPIHandler
// i.e. convert query params, path parms, & body to golang args

type httpHandler struct {
	a APIHandler
}

func apiHandlerToHTTPAPIHandler(a APIHandler) HTTPAPIHandler {
	return &httpHandler{
		a: a,
	}
}
func (h *httpHandler) GetStatus(w http.ResponseWriter, req *http.Request) {
	r, err := h.a.GetStatus(req.Context())
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, r)
}
func (h *httpHandler) ListDocuments(w http.ResponseWriter, req *http.Request) {
	pageToken := req.URL.Query().Get("page_token")
	pageSize, _ := strconv.Atoi(req.URL.Query().Get("page_size"))
	r, err := h.a.ListDocuments(req.Context(), &model.ListDocumentsQueryParams{
		PageToken: pageToken,
		PageSize:  pageSize,
	})
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, r)
}
func (h *httpHandler) CreateDocument(w http.ResponseWriter, req *http.Request) {
	var d model.Document
	if err := json.NewDecoder(req.Body).Decode(&d); err != nil {
		SendError(w, err)
		return
	}
	r, err := h.a.CreateDocument(req.Context(), &d)
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, r)
}
func (h *httpHandler) GetDocument(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	documentID, ok := vars["documentID"]
	if !ok {
		SendError(w, NewHTTPErrorFromString(http.StatusBadRequest, "invalid `documentID` path parameter"))
		return
	}
	documentID, err := url.QueryUnescape(documentID)
	if err != nil {
		SendError(w, NewHTTPErrorFromString(http.StatusBadRequest, "bad `documentID` path parameter"))
		return
	}
	r, err := h.a.GetDocument(req.Context(), &model.GetDocumentPathParams{DocumentID: documentID})
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, r)
}
func (h *httpHandler) PutDocument(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	documentID, ok := vars["documentID"]
	if !ok {
		SendError(w, NewHTTPErrorFromString(http.StatusBadRequest, "invalid `documentID` path parameter"))
		return
	}
	documentID, err := url.QueryUnescape(documentID)
	if err != nil {
		SendError(w, NewHTTPErrorFromString(http.StatusBadRequest, "bad `documentID` path parameter"))
		return
	}
	var d model.Document
	if err := json.NewDecoder(req.Body).Decode(&d); err != nil {
		SendError(w, err)
		return
	}
	r, err := h.a.PutDocument(req.Context(), &model.PutDocumentPathParams{DocumentID: documentID}, &d)
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, r)
}
func (h *httpHandler) DeleteDocument(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	documentID, ok := vars["documentID"]
	if !ok {
		SendError(w, NewHTTPErrorFromString(http.StatusBadRequest, "invalid `documentID` path parameter"))
		return
	}
	documentID, err := url.QueryUnescape(documentID)
	if err != nil {
		SendError(w, NewHTTPErrorFromString(http.StatusBadRequest, "bad `documentID` path parameter"))
		return
	}
	err = h.a.DeleteDocument(req.Context(), &model.DeleteDocumentPathParams{DocumentID: documentID})
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, struct{}{})
}
func (h *httpHandler) QueryDocuments(w http.ResponseWriter, req *http.Request) {
	var b model.QueryDocumentsRequest
	if err := json.NewDecoder(req.Body).Decode(&b); err != nil {
		SendError(w, err)
		return
	}
	r, err := h.a.QueryDocuments(req.Context(), &b)
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, r)
}
func (h *httpHandler) QueueJob(w http.ResponseWriter, req *http.Request) {
	var j model.Job
	if err := json.NewDecoder(req.Body).Decode(&j); err != nil {
		SendError(w, err)
		return
	}
	r, err := h.a.QueueJob(req.Context(), &j)
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, r)
}
func (h *httpHandler) GetJob(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jobID, ok := vars["jobID"]
	if !ok {
		SendError(w, NewHTTPErrorFromString(http.StatusBadRequest, "invalid `jobID` path parameter"))
		return
	}
	r, err := h.a.GetJob(req.Context(), &model.GetJobPathParams{JobID: jobID})
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, r)
}
func (h *httpHandler) ListJobs(w http.ResponseWriter, req *http.Request) {
	pageToken := req.URL.Query().Get("page_token")
	pageSize, _ := strconv.Atoi(req.URL.Query().Get("page_size"))
	r, err := h.a.ListJobs(req.Context(), &model.ListJobsQueryParams{
		PageToken: pageToken,
		PageSize:  pageSize,
	})
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, r)
}
func (h *httpHandler) DeleteJob(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jobID, ok := vars["jobID"]
	if !ok {
		SendError(w, NewHTTPErrorFromString(http.StatusBadRequest, "invalid `jobID` path parameter"))
		return
	}
	err := h.a.DeleteJob(req.Context(), &model.DeleteJobPathParams{JobID: jobID})
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, struct{}{})
}
func (h *httpHandler) ClaimSomeJob(w http.ResponseWriter, req *http.Request) {
	var body model.ClaimJobRequest
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		SendError(w, err)
		return
	}
	r, err := h.a.ClaimSomeJob(req.Context(), &body)
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, r)
}
func (h *httpHandler) ClaimJob(w http.ResponseWriter, req *http.Request) {
	var body model.ClaimJobRequest
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		SendError(w, err)
		return
	}
	vars := mux.Vars(req)
	jobID, ok := vars["jobID"]
	if !ok {
		SendError(w, NewHTTPErrorFromString(http.StatusBadRequest, "invalid `jobID` path parameter"))
		return
	}
	r, err := h.a.ClaimJob(req.Context(), &model.ClaimJobPathParams{JobID: jobID})
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, r)
}
func (h *httpHandler) ReleaseJob(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jobID, ok := vars["jobID"]
	if !ok {
		SendError(w, NewHTTPErrorFromString(http.StatusBadRequest, "invalid `jobID` path parameter"))
		return
	}
	r, err := h.a.ReleaseJob(req.Context(), &model.ReleaseJobPathParams{JobID: jobID})
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, r)
}
func (h *httpHandler) SetJobSuccess(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jobID, ok := vars["jobID"]
	if !ok {
		SendError(w, NewHTTPErrorFromString(http.StatusBadRequest, "invalid `jobID` path parameter"))
		return
	}
	var j model.Job
	if err := json.NewDecoder(req.Body).Decode(&j); err != nil {
		SendError(w, err)
		return
	}
	r, err := h.a.SetJobSuccess(req.Context(), &model.SetJobSuccessPathParams{JobID: jobID}, &j)
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, r)
}
func (h *httpHandler) SetJobError(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jobID, ok := vars["jobID"]
	if !ok {
		SendError(w, NewHTTPErrorFromString(http.StatusBadRequest, "invalid `jobID` path parameter"))
		return
	}
	var j model.Job
	if err := json.NewDecoder(req.Body).Decode(&j); err != nil {
		SendError(w, err)
		return
	}
	r, err := h.a.SetJobError(req.Context(), &model.SetJobErrorPathParams{JobID: jobID}, &j)
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, r)
}

func (h *httpHandler) ListSecrets(w http.ResponseWriter, req *http.Request) {
	pageToken := req.URL.Query().Get("page_token")
	pageSize, _ := strconv.Atoi(req.URL.Query().Get("page_size"))
	v, err := h.a.ListSecrets(req.Context(), &model.ListSecretsRequest{
		PageToken: pageToken,
		PageSize:  pageSize,
	})
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, v)
}
func (h *httpHandler) GetSecret(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key, ok := vars["key"]
	if !ok {
		SendError(w, NewHTTPErrorFromString(http.StatusBadRequest, "invalid `key` path parameter"))
		return
	}
	v, err := h.a.GetSecret(req.Context(), key)
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, v)
}
func (h *httpHandler) SetSecret(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key, ok := vars["key"]
	if !ok {
		SendError(w, NewHTTPErrorFromString(http.StatusBadRequest, "invalid `key` path parameter"))
		return
	}
	v, err := h.a.SetSecret(req.Context(), key)
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, v)
}
func (h *httpHandler) ListOAuthProviders(w http.ResponseWriter, req *http.Request) {
	pageToken := req.URL.Query().Get("page_token")
	pageSize, _ := strconv.Atoi(req.URL.Query().Get("page_size"))
	v, err := h.a.ListOAuthProviders(req.Context(), &model.ListOAuthProvidersRequest{
		PageToken: pageToken,
		PageSize:  pageSize,
	})
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, v)
}
func (h *httpHandler) GetOAuthAuthorizationURL(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		SendError(w, NewHTTPErrorFromString(http.StatusBadRequest, "invalid `name` path parameter"))
		return
	}
	v, err := h.a.GetOAuthAuthorizationURL(req.Context(), name)
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, v)
}
func (h *httpHandler) GetOAuthToken(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		SendError(w, NewHTTPErrorFromString(http.StatusBadRequest, "invalid `name` path parameter"))
		return
	}
	b := model.GetOAuthTokenRequest{}
	if err := json.NewDecoder(req.Body).Decode(&b); err != nil {
		SendError(w, err)
		return
	}
	v, err := h.a.GetOAuthToken(req.Context(), name, &b)
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, v)
}
func (h *httpHandler) GetOAuthConfig(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		SendError(w, NewHTTPErrorFromString(http.StatusBadRequest, "invalid `name` path parameter"))
		return
	}
	v, err := h.a.GetOAuthConfig(req.Context(), name)
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, v)
}
func (h *httpHandler) SetOAuthConfig(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name, ok := vars["name"]
	if !ok {
		SendError(w, NewHTTPErrorFromString(http.StatusBadRequest, "invalid `name` path parameter"))
		return
	}
	v, err := h.a.SetOAuthConfig(req.Context(), name)
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, v)
}
