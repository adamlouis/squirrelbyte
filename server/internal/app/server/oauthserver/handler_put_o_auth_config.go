package oauthserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/oauthmodel"
)

func (h *hdl) PutOAuthConfig(ctx context.Context, pathParams *oauthmodel.PutOAuthConfigPathParams, body *oauthmodel.Config) (*oauthmodel.Config, int, error) {
	return nil, http.StatusInternalServerError, fmt.Errorf("unimplemented")
}
