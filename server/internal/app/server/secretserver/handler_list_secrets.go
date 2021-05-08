package secretserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/secretmodel"
)

func (h *hdl) ListSecrets(ctx context.Context, queryParams *secretmodel.ListSecretsRequest) (*secretmodel.ListSecretsResponse, int, error) {
	return nil, http.StatusInternalServerError, fmt.Errorf("unimplemented")
}
