package serverhandler

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

func (a *apiHandler) GetOAuthConfig(ctx context.Context, provider string) (*model.Config, error) {
	r, err := a.GetRepositories()
	if err != nil {
		return nil, err
	}
	defer r.Rollback()
	return r.OauthManager.GetConfig(ctx, provider)
}
