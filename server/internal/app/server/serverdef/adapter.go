package serverdef

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

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
	r, err := h.a.ListDocuments(req.Context(), &ListDocumentsQueryParams{
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
	var d Document
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
	r, err := h.a.GetDocument(req.Context(), &GetDocumentPathParams{DocumentID: documentID})
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
	var d Document
	if err := json.NewDecoder(req.Body).Decode(&d); err != nil {
		SendError(w, err)
		return
	}
	r, err := h.a.PutDocument(req.Context(), &PutDocumentPathParams{DocumentID: documentID}, &d)
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
	err = h.a.DeleteDocument(req.Context(), &DeleteDocumentPathParams{DocumentID: documentID})
	if err != nil {
		SendError(w, err)
		return
	}
	SendOK(w, struct{}{})
}
func (h *httpHandler) QueryDocuments(w http.ResponseWriter, req *http.Request) {
	var b QueryDocumentsRequest
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
