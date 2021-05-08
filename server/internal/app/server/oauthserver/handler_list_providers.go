package oauthserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/oauthmodel"
)

func (h *hdl) ListProviders(ctx context.Context, queryParams *oauthmodel.ListOAuthProvidersRequest) (*oauthmodel.ListOAuthProvidersResponse, int, error) {
	return nil, http.StatusInternalServerError, fmt.Errorf("unimplemented")
}
