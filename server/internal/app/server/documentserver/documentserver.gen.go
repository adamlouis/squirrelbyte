// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: --component server --config ../../../../config/api.document.yml --package documentserver --out-dir . --out ./documentserver.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel
package documentserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type HTTPHandler interface {
	ListDocuments(w http.ResponseWriter, req *http.Request)
	PostDocument(w http.ResponseWriter, req *http.Request)
	GetDocument(w http.ResponseWriter, req *http.Request)
	PutDocument(w http.ResponseWriter, req *http.Request)
	DeleteDocument(w http.ResponseWriter, req *http.Request)
	QueryDocuments(w http.ResponseWriter, req *http.Request)
}
type APIHandler interface {
	ListDocuments(ctx context.Context, queryParams *documentmodel.ListDocumentsQueryParams) (*documentmodel.ListDocumentsResponse, error)
	PostDocument(ctx context.Context, body *documentmodel.Document) (*documentmodel.Document, error)
	GetDocument(ctx context.Context, pathParams *documentmodel.GetDocumentPathParams) (*documentmodel.Document, error)
	PutDocument(ctx context.Context, pathParams *documentmodel.PutDocumentPathParams, body *documentmodel.Document) (*documentmodel.Document, error)
	DeleteDocument(ctx context.Context, pathParams *documentmodel.DeleteDocumentPathParams) error
	QueryDocuments(ctx context.Context, body *documentmodel.QueryDocumentsRequest) (*documentmodel.QueryDocumentsResponse, error)
}

func RegisterRouter(apiHandler APIHandler, r *mux.Router, c ErrorCoder) {
	h := apiHandlerToHTTPHandler(apiHandler, c)
	r.Handle("/documents", http.HandlerFunc(h.ListDocuments)).Methods(http.MethodGet)
	r.Handle("/documents", http.HandlerFunc(h.PostDocument)).Methods(http.MethodPost)
	r.Handle("/documents/{documentID}", http.HandlerFunc(h.GetDocument)).Methods(http.MethodGet)
	r.Handle("/documents/{documentID}", http.HandlerFunc(h.PutDocument)).Methods(http.MethodPut)
	r.Handle("/documents/{documentID}", http.HandlerFunc(h.DeleteDocument)).Methods(http.MethodDelete)
	r.Handle("/documents:query", http.HandlerFunc(h.QueryDocuments)).Methods(http.MethodPost)
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

func (h *httpHandler) ListDocuments(w http.ResponseWriter, req *http.Request) {
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
	queryParams := documentmodel.ListDocumentsQueryParams{
		PageToken: pageTokenQueryParam,
		PageSize:  pageSizeQueryParam,
	}
	r, err := h.apiHandler.ListDocuments(req.Context(), &queryParams)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) PostDocument(w http.ResponseWriter, req *http.Request) {
	var requestBody documentmodel.Document
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		sendErrorWithCode(w, http.StatusBadRequest, err)
		return
	}
	r, err := h.apiHandler.PostDocument(req.Context(), &requestBody)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) GetDocument(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	documentID, ok := vars["documentID"]
	if !ok {
		sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid documentID path parameter"))
		return
	}
	pathParams := documentmodel.GetDocumentPathParams{
		DocumentID: documentID,
	}
	r, err := h.apiHandler.GetDocument(req.Context(), &pathParams)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) PutDocument(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	documentID, ok := vars["documentID"]
	if !ok {
		sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid documentID path parameter"))
		return
	}
	pathParams := documentmodel.PutDocumentPathParams{
		DocumentID: documentID,
	}
	var requestBody documentmodel.Document
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		sendErrorWithCode(w, http.StatusBadRequest, err)
		return
	}
	r, err := h.apiHandler.PutDocument(req.Context(), &pathParams, &requestBody)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
func (h *httpHandler) DeleteDocument(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	documentID, ok := vars["documentID"]
	if !ok {
		sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid documentID path parameter"))
		return
	}
	pathParams := documentmodel.DeleteDocumentPathParams{
		DocumentID: documentID,
	}
	err := h.apiHandler.DeleteDocument(req.Context(), &pathParams)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, struct{}{})
}
func (h *httpHandler) QueryDocuments(w http.ResponseWriter, req *http.Request) {
	var requestBody documentmodel.QueryDocumentsRequest
	if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
		sendErrorWithCode(w, http.StatusBadRequest, err)
		return
	}
	r, err := h.apiHandler.QueryDocuments(req.Context(), &requestBody)
	if err != nil {
		h.sendError(w, err)
		return
	}
	sendOK(w, r)
}
