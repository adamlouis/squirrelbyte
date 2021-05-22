package kv

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/kvmodel"
)

type Repository interface {
	Put(ctx context.Context, s *kvmodel.KV) (*kvmodel.KV, error)
	Get(ctx context.Context, key string) (*kvmodel.KV, error)
	List(ctx context.Context, args *kvmodel.ListKVsRequest) (*kvmodel.ListKVsResponse, error)
	Delete(ctx context.Context, key string) error
}
