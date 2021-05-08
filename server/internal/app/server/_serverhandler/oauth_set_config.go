package serverhandler

import (
	"context"
	"fmt"

	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

func (a *apiHandler) SetOAuthConfig(ctx context.Context, provider string, config *model.Config) (*model.Config, error) {
	r, err := a.GetRepositories()
	if err != nil {
		return nil, err
	}
	defer r.Rollback()

	c, err := r.OauthManager.SetConfig(ctx, provider, config)
	if err != nil {
		return nil, fmt.Errorf("config not found for provider %s", provider)
	}

	if err = r.Commit(); err != nil {
		return nil, err
	}

	return c, nil
}
