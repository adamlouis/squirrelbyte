package secretserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/secretmodel"
)

func (h *hdl) GetSecret(ctx context.Context, pathParams *secretmodel.GetSecretPathParams) (*secretmodel.Secret, int, error) {
	return nil, http.StatusInternalServerError, fmt.Errorf("unimplemented")
}
