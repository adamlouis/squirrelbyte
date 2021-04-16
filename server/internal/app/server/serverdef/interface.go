package serverdef

import (
	"context"
	"net/http"
)

// TODO: generate all of `serverdef` package from conf / openapi declaration

// HTTPAPIHandler is an implementation of the API using the http.HandlerFunc signature
type HTTPAPIHandler interface {
	GetStatus(w http.ResponseWriter, req *http.Request)
	ListDocuments(w http.ResponseWriter, req *http.Request)
	CreateDocument(w http.ResponseWriter, req *http.Request)
	GetDocument(w http.ResponseWriter, req *http.Request)
	PutDocument(w http.ResponseWriter, req *http.Request)
	DeleteDocument(w http.ResponseWriter, req *http.Request)
	QueryDocuments(w http.ResponseWriter, req *http.Request)
}

// APIHandler is an implementation of the API using typed golang structs
type APIHandler interface {
	GetStatus(ctx context.Context) (*Status, error)
	ListDocuments(ctx context.Context, queryParams *ListDocumentsQueryParams) (*ListDocumentsResponse, error)
	CreateDocument(ctx context.Context, document *Document) (*Document, error)
	GetDocument(ctx context.Context, pathParams *GetDocumentPathParams) (*Document, error)
	PutDocument(ctx context.Context, pathParams *PutDocumentPathParams, document *Document) (*Document, error)
	DeleteDocument(ctx context.Context, pathParams *DeleteDocumentPathParams) error
	QueryDocuments(ctx context.Context, body *QueryDocumentsRequest) (*QueryDocumentsResponse, error)
}
