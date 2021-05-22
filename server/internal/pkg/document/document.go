package document

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
)

// Repository is the interface for storing & accessing documents
type Repository interface {
	Put(ctx context.Context, d *documentmodel.Document) (*documentmodel.Document, error)
	Get(ctx context.Context, documentID string) (*documentmodel.Document, error)
	List(ctx context.Context, args *documentmodel.ListDocumentsQueryParams) (*documentmodel.ListDocumentsResponse, error)
	Delete(ctx context.Context, documentID string) error
	Query(ctx context.Context, q *documentmodel.QueryDocumentsRequest) (*documentmodel.QueryDocumentsResponse, error)
}
