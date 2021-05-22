package oauth

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/oauthmodel"
	"golang.org/x/oauth2"
)

type Repository interface {
	PutConfig(ctx context.Context, c *oauthmodel.Config) (*oauthmodel.Config, error)
	GetConfig(ctx context.Context, name string) (*oauthmodel.Config, error)
	DeleteConfig(ctx context.Context, name string) error
	ListProviders(ctx context.Context, b *oauthmodel.ListOAuthProvidersRequest) (*oauthmodel.ListOAuthProvidersResponse, error)
}

func ToLibConfig(c *oauthmodel.Config) *oauth2.Config {
	if c == nil {
		return nil
	}
	r := &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		RedirectURL:  c.RedirectURL,
		Scopes:       c.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  c.AuthURL,
			TokenURL: c.TokenURL,
		},
	}
	return r
}
