package oauthserver

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/oauth"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/oauthmodel"
	"golang.org/x/oauth2"
)

func (h *hdl) GetOAuthAuthorizationURL(ctx context.Context, pathParams *oauthmodel.GetOAuthAuthorizationURLPathParams) (*oauthmodel.GetOAuthAuthorizationURLResponse, error) {
	repo, _, rollback, err := h.GetRepository()
	if err != nil {
		return nil, err
	}
	defer rollback() //nolint

	c, err := repo.GetConfig(ctx, pathParams.Name)
	if err != nil {
		return nil, fmt.Errorf("config not found for provider %s", pathParams.Name)
	}

	l := oauth.ToLibConfig(c)

	opts := make([]oauth2.AuthCodeOption, len(c.AuthURLParams))
	i := 0
	for k, v := range c.AuthURLParams {
		opts[i] = oauth2.SetAuthURLParam(k, v)
		i += 1
	}

	return &oauthmodel.GetOAuthAuthorizationURLResponse{
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
