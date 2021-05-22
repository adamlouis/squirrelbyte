package oauthserver

import (
	"context"
	"fmt"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/oauthmodel"
)

func (h *hdl) PutOAuthConfig(ctx context.Context, pathParams *oauthmodel.PutOAuthConfigPathParams, body *oauthmodel.Config) (*oauthmodel.Config, error) {
	if pathParams.Name != body.Name {
		return nil, fmt.Errorf("name in path does not match name in request body")
	}

	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return nil, err
	}
	defer rollback() //nolint

	c, err := repo.PutConfig(ctx, body)
	if err != nil {
		return nil, err
	}

	if err = commit(); err != nil {
		return nil, err
	}

	return c, nil
}
