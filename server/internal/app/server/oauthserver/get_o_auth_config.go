package oauthserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/oauthmodel"
)

func (h *hdl) GetOAuthConfig(ctx context.Context, pathParams *oauthmodel.GetOAuthConfigPathParams, body *oauthmodel.Config) (*oauthmodel.Config, int, error) {
	return nil, http.StatusInternalServerError, fmt.Errorf("unimplemented")
}
