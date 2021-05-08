package serverhandler

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

func (a *apiHandler) ListOAuthProviders(ctx context.Context, body *model.ListOAuthProvidersRequest) (*model.ListOAuthProvidersResponse, error) {
	r, err := a.GetRepositories()
	if err != nil {
		return nil, err
	}
	defer r.Rollback()
	p, err := r.OauthManager.ListProviders(ctx, body)
	if err != nil {
		return nil, err
	}

	ps := []*model.Provider{}
	for k, v := range p {
		ps = append(ps, &model.Provider{
			Name:   k,
			Config: v,
		})
	}

	return &model.ListOAuthProvidersResponse{
		Providers: ps,
	}, nil
}
