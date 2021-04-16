package document

import (
	"context"
	"time"
)

// Document is the document resource
type Document struct {
	ID        string
	Header    []byte
	Body      []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Repository is the interface for storing & accessing documents
type Repository interface {
	Init(ctx context.Context) error
	Put(ctx context.Context, d *Document) (*Document, error)
	Get(ctx context.Context, documentID string) (*Document, error)
	List(ctx context.Context, args *ListDocumentArgs) (*ListDocumentResult, error)
	Delete(ctx context.Context, documentID string) error
	Query(ctx context.Context, q *Query) (*QueryResult, error)
}

// PageArgs are the arguments getting a page
type PageArgs struct {
	PageSize  int
	PageToken string
}

// PageResult are the values getting the next page
type PageResult struct {
	NextPageToken string
}

// ListDocumentArgs are the args for listing documents
type ListDocumentArgs struct {
	PageArgs
}

// ListDocumentResult is the result of listing documents
type ListDocumentResult struct {
	PageResult
	Documents []*Document
}

// Query is the document query structure
type Query struct {
	PageToken string
	Select    []interface{}
	GroupBy   []interface{}
	OrderBy   []interface{}
	Where     interface{}
	Limit     int
}

// QueryResult is the result of querying documents
type QueryResult struct {
	PageResult
	Result []interface{}
}

// SearchPathsQuery is query for searching paths
type SearchPathsQuery struct {
	PageArgs
	Query string
}

// SearchPathsResult is result of searching paths
type SearchPathsResult struct {
	PageResult
	Paths []string
}
