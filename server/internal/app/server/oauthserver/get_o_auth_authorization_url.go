package oauthserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/oauthmodel"
)

func (h *hdl) GetOAuthAuthorizationURL(ctx context.Context, pathParams *oauthmodel.GetOAuthAuthorizationURLPathParams) (*oauthmodel.GetOAuthAuthorizationURLResponse, int, error) {
	return nil, http.StatusInternalServerError, fmt.Errorf("unimplemented")
}
