package serverhandler

import (
	"context"
	"time"

	"github.com/adamlouis/squirrelbyte/server/internal/app/server/serverdef"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
)

func (a *apiHandler) GetStatus(ctx context.Context) (*serverdef.Status, error) {
	return &serverdef.Status{
		Status:    "OK",
		Timestamp: present.ToAPITime(time.Now()),
	}, nil
}
