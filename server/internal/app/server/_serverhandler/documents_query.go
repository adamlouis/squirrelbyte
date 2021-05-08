package serverhandler

import (
	"context"
	"time"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/document"
	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

func (a *apiHandler) QueryDocuments(ctx context.Context, body *model.QueryDocumentsRequest) (*model.QueryDocumentsResponse, error) {
	start := time.Now()
	repos, err := a.GetRepositories()
	if err != nil {
		return nil, err
	}
	defer repos.Rollback() //nolint

	r, err := repos.Document.Query(ctx, &document.Query{
		Select:    body.Select,
		Where:     body.Where,
		GroupBy:   body.GroupBy,
		OrderBy:   body.OrderBy,
		Limit:     body.Limit,
		PageToken: body.PageToken,
	})
	if err != nil {
		return nil, err
	}

	return &model.QueryDocumentsResponse{
		Result:        r.Result,
		NextPageToken: r.NextPageToken,
		Insights: map[string]interface{}{
			"duration_ms": time.Since(start) / time.Millisecond,
		},
	}, nil
}
