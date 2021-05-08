package serverhandler

import (
	"context"
	"fmt"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/oauth"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

func (a *apiHandler) GetOAuthToken(ctx context.Context, provider string, body *model.GetOAuthTokenRequest) (*model.Token, error) {
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

	tk, err := l.Exchange(ctx, body.Code)
	if err != nil {
		return nil, err
	}

	return &model.Token{
		AccessToken:  tk.AccessToken,
		TokenType:    tk.TokenType,
		RefreshToken: tk.RefreshToken,
		Expiry:       present.ToAPITime(tk.Expiry),
	}, err
}
