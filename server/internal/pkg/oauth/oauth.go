package oauth

import (
	"context"
	"fmt"
	"os"

	"github.com/adamlouis/squirrelbyte/server/pkg/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

type Manager interface {
	SetConfig(ctx context.Context, key string, c *model.Config) (*model.Config, error)
	GetConfig(ctx context.Context, key string) (*model.Config, error)
	// TODO: wrong interface
	ListProviders(ctx context.Context, b *model.ListOAuthProvidersRequest) (map[string]*model.Config, error)
}

func NewManger() Manager {
	return &mgr{
		configs: map[string]*model.Config{
			"strava": {
				ClientID:     os.Getenv("STRAVA_CLIENT_ID"),
				ClientSecret: os.Getenv("STRAVA_CLIENT_SECRET"),
				Endpoint: &model.Endpoint{
					AuthURL:  "https://www.strava.com/api/v3/oauth/authorize",
					TokenURL: "https://www.strava.com/api/v3/oauth/token",
				},
				RedirectURL: "http://localhost:9921/oauth/providers/strava/token",
				Scopes:      []string{"activity:read_all"},
			},
			"spotify": {
				ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
				ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
				Endpoint: &model.Endpoint{
					AuthURL:  spotify.Endpoint.AuthURL,
					TokenURL: spotify.Endpoint.TokenURL,
				},
				RedirectURL: "http://localhost:9921/oauth/providers/spotify/token",
				Scopes: []string{
					// https://developer.spotify.com/documentation/general/guides/scopes/
					"user-read-recently-played",
					"user-top-read",
					"playlist-read-private",
					"playlist-read-collaborative",
					"user-follow-read",
					"user-library-read",
				},
			},
		},
	}
}

type mgr struct {
	configs map[string]*model.Config
}

func (m *mgr) SetConfig(ctx context.Context, key string, c *model.Config) (*model.Config, error) {
	m.configs[key] = c
	return m.GetConfig(ctx, key)
}

func (m *mgr) GetConfig(ctx context.Context, key string) (*model.Config, error) {
	v, ok := m.configs[key]
	if !ok {
		return nil, fmt.Errorf("config for `%s` not found", key)
	}
	return v, nil
}

func (m *mgr) ListProviders(ctx context.Context, b *model.ListOAuthProvidersRequest) (map[string]*model.Config, error) {
	return m.configs, nil
}

func ToLibConfig(c *model.Config) *oauth2.Config {
	if c == nil {
		return nil
	}
	r := &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		RedirectURL:  c.RedirectURL,
		Scopes:       c.Scopes,
	}

	if c.Endpoint != nil {
		r.Endpoint.AuthURL = c.Endpoint.AuthURL
		r.Endpoint.TokenURL = c.Endpoint.TokenURL
	}

	return r
}
