package document

import (
	"context"
	"time"
)

type Document struct {
	ID        string
	Header    []byte
	Body      []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Repository interface {
	Init(ctx context.Context) error
	Put(ctx context.Context, d *Document) (*Document, error)
	Get(ctx context.Context, documentID string) (*Document, error)
	List(ctx context.Context, args *ListDocumentArgs) (*ListDocumentResult, error)
	Delete(ctx context.Context, documentID string) error
	Search(ctx context.Context, q *SearchDocumentsQuery) (*SearchDocumentsResult, error)
}

type PageArgs struct {
	PageSize  int
	PageToken string
}

type PageResult struct {
	NextPageToken string
}

type ListDocumentArgs struct {
	PageArgs
}

type ListDocumentResult struct {
	PageResult
	Documents []*Document
}

type SearchDocumentsQuery struct {
	PageToken string
	Select    []interface{}
	GroupBy   []interface{}
	OrderBy   []interface{}
	Where     interface{}
	Limit     int
}

type SearchDocumentsResult struct {
	PageResult
	Result []interface{}
}

type SearchPathsQuery struct {
	PageArgs
	Query string
}

type SearchPathsResult struct {
	PageResult
	Paths []string
}
