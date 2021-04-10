package serverdef

import (
	"net/http"

	"github.com/gorilla/mux"
)

// TODO: generate all of `serverdef` package from conf / openapi declaration

// RegisterRouter registers the router with the api handler
func RegisterRouter(a APIHandler, r *mux.Router) {
	h := apiHandlerToHTTPAPIHandler(a)
	r.Handle("/", http.HandlerFunc(h.GetStatus)).Methods(http.MethodGet)
	r.Handle("/status", http.HandlerFunc(h.GetStatus)).Methods(http.MethodGet)
	r.Handle("/documents", http.HandlerFunc(h.ListDocuments)).Methods(http.MethodGet)
	r.Handle("/documents", http.HandlerFunc(h.CreateDocument)).Methods(http.MethodPost)
	r.Handle("/documents/{documentID}", http.HandlerFunc(h.GetDocument)).Methods(http.MethodGet)
	r.Handle("/documents/{documentID}", http.HandlerFunc(h.PutDocument)).Methods(http.MethodPut)
	r.Handle("/documents/{documentID}", http.HandlerFunc(h.DeleteDocument)).Methods(http.MethodDelete)
	r.Handle("/documents:search", http.HandlerFunc(h.SearchDocuments)).Methods(http.MethodPost)
}
