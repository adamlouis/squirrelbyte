package serverhandler

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/oauth"
	"github.com/adamlouis/squirrelbyte/server/pkg/model"
	"golang.org/x/oauth2"
)

func (a *apiHandler) GetOAuthAuthorizationURL(ctx context.Context, provider string) (*model.GetOAuthAuthorizationURLResponse, error) {
	r, err := a.GetRepositories()
	if err != nil {
		return nil, err
	}
	defer r.Rollback()

	c, err := r.OauthManager.GetConfig(ctx, provider)
	if err != nil {
		return nil, fmt.Errorf("config not found for provider %s", provider)
	}

	l := oauth.ToLibConfig(c)

	opts := make([]oauth2.AuthCodeOption, len(c.AdditionalAuthorizationURLParams))
	for i, o := range c.AdditionalAuthorizationURLParams {
		opts[i] = oauth2.SetAuthURLParam(o.Key, o.Value)
	}

	return &model.GetOAuthAuthorizationURLResponse{
		URL: l.AuthCodeURL(getRandomString(16), opts...),
	}, nil
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
