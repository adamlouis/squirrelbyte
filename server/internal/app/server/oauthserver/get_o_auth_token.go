package oauthserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/oauthmodel"
)

func (h *hdl) GetOAuthToken(ctx context.Context, pathParams *oauthmodel.GetOAuthTokenPathParams, body *oauthmodel.GetOAuthTokenRequest) (*oauthmodel.Token, int, error) {
	return nil, http.StatusInternalServerError, fmt.Errorf("unimplemented")
}
