package serverhandler

import (
	"context"
	"time"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

func (a *apiHandler) GetStatus(ctx context.Context) (*model.Status, error) {
	return &model.Status{
		Status:    "OK",
		Timestamp: present.ToAPITime(time.Now()),
	}, nil
}
