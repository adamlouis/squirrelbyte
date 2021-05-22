package oauthserver

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/oauth"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/oauthmodel"
)

func (h *hdl) GetOAuthToken(ctx context.Context, pathParams *oauthmodel.GetOAuthTokenPathParams, body *oauthmodel.GetOAuthTokenRequest) (*oauthmodel.Token, error) {
	repo, _, rollback, err := h.GetRepository()
	if err != nil {
		return nil, err
	}
	defer rollback() //nolint

	c, err := repo.GetConfig(ctx, pathParams.Name)
	if err != nil {
		return nil, err
	}

	l := oauth.ToLibConfig(c)

	tk, err := l.Exchange(ctx, body.Code)
	if err != nil {
		return nil, err
	}

	return &oauthmodel.Token{
		AccessToken:  tk.AccessToken,
		TokenType:    tk.TokenType,
		RefreshToken: tk.RefreshToken,
		Expiry:       present.ToAPITime(tk.Expiry),
	}, nil
}
