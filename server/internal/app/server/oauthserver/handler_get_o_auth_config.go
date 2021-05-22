package oauthserver

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/oauthmodel"
)

func (h *hdl) GetOAuthConfig(ctx context.Context, pathParams *oauthmodel.GetOAuthConfigPathParams) (*oauthmodel.Config, error) {
	repo, _, rollback, err := h.GetRepository()
	if err != nil {
		return nil, err
	}
	defer rollback() //nolint

	return repo.GetConfig(ctx, pathParams.Name)
}
