package serverdef

import (
	"context"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

// TODO: generate all of `serverdef` package from conf / openapi declaration

// HTTPAPIHandler is an implementation of the API using the http.HandlerFunc signature
type HTTPAPIHandler interface {
	GetStatus(w http.ResponseWriter, req *http.Request)
	// documents
	ListDocuments(w http.ResponseWriter, req *http.Request)
	CreateDocument(w http.ResponseWriter, req *http.Request)
	GetDocument(w http.ResponseWriter, req *http.Request)
	PutDocument(w http.ResponseWriter, req *http.Request)
	DeleteDocument(w http.ResponseWriter, req *http.Request)
	QueryDocuments(w http.ResponseWriter, req *http.Request)
	// jobs
	GetJob(w http.ResponseWriter, req *http.Request)
	ListJobs(w http.ResponseWriter, req *http.Request)
	QueueJob(w http.ResponseWriter, req *http.Request)
	DeleteJob(w http.ResponseWriter, req *http.Request)
	ClaimJob(w http.ResponseWriter, req *http.Request)
	ClaimSomeJob(w http.ResponseWriter, req *http.Request)
	ReleaseJob(w http.ResponseWriter, req *http.Request)
	SetJobSuccess(w http.ResponseWriter, req *http.Request)
	SetJobError(w http.ResponseWriter, req *http.Request)
}

// APIHandler is an implementation of the API using typed golang structs
type APIHandler interface {
	GetStatus(ctx context.Context) (*model.Status, error)
	ListDocuments(ctx context.Context, queryParams *model.ListDocumentsQueryParams) (*model.ListDocumentsResponse, error)
	CreateDocument(ctx context.Context, document *model.Document) (*model.Document, error)
	GetDocument(ctx context.Context, pathParams *model.GetDocumentPathParams) (*model.Document, error)
	PutDocument(ctx context.Context, pathParams *model.PutDocumentPathParams, document *model.Document) (*model.Document, error)
	DeleteDocument(ctx context.Context, pathParams *model.DeleteDocumentPathParams) error
	QueryDocuments(ctx context.Context, body *model.QueryDocumentsRequest) (*model.QueryDocumentsResponse, error)
	GetJob(ctx context.Context, pathParams *model.GetJobPathParams) (*model.Job, error)
	ListJobs(ctx context.Context, queryParams *model.ListJobsQueryParams) (*model.ListJobsResponse, error)
	QueueJob(ctx context.Context, body *model.Job) (*model.Job, error)
	DeleteJob(ctx context.Context, pathParams *model.DeleteJobPathParams) error
	ClaimJob(ctx context.Context, pathParams *model.ClaimJobPathParams) (*model.Job, error)
	ClaimSomeJob(ctx context.Context, body *model.ClaimJobRequest) (*model.Job, error)
	ReleaseJob(ctx context.Context, pathParams *model.ReleaseJobPathParams) (*model.Job, error)
	SetJobSuccess(ctx context.Context, pathParams *model.SetJobSuccessPathParams, job *model.Job) (*model.Job, error)
	SetJobError(ctx context.Context, pathParams *model.SetJobErrorPathParams, job *model.Job) (*model.Job, error)
}
