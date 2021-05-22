package oauthserver

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/oauthmodel"
)

func (h *hdl) DeleteOAuthConfig(ctx context.Context, pathParams *oauthmodel.DeleteOAuthConfigPathParams) error {
	repo, commit, rollback, err := h.GetRepository()
	if err != nil {
		return err
	}
	defer rollback() //nolint

	if err = repo.DeleteConfig(ctx, pathParams.Name); err != nil {
		return err
	}

	return commit()
}
