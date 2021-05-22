package oauthsqlite3

import (
	"context"
	"database/sql"
	"embed"
	"encoding/json"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/crudutil"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/errtype"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/oauth"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/sqlite3util"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/oauthmodel"
	"github.com/jmoiron/sqlx"
)

func NewRepository(db sqlx.Ext) oauth.Repository {
	return &repo{
		db: db,
	}
}

type repo struct {
	db sqlx.Ext
}

//go:embed migration/*.sql
var MigrationFS embed.FS

type configRow struct {
	Name          string `db:"name"`
	ClientID      string `db:"client_id"`
	ClientSecret  string `db:"client_secret"`
	AuthURL       string `db:"auth_url"`
	TokenURL      string `db:"token_url"`
	RedirectURL   string `db:"redirect_url"`
	Scopes        []byte `db:"scopes"`
	AuthURLParams []byte `db:"auth_url_params"`
	CreatedAt     string `db:"created_at"`
	UpdatedAt     string `db:"updated_at"`
}

func (r *repo) PutConfig(ctx context.Context, c *oauthmodel.Config) (*oauthmodel.Config, error) {
	scopes, err := json.Marshal(c.Scopes)
	if err != nil {
		return nil, err
	}

	authURLParams, err := json.Marshal(c.AuthURLParams)
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(`
			INSERT INTO
				config(
					name,
					client_id,
					client_secret,
					auth_url,
					auth_url_params,
					token_url,
					redirect_url,
					scopes
				)
			VALUES
				(
					?,
					?,
					?,
					?,
					?,
					?,
					?,
					?
				)
			ON CONFLICT (name)
			DO UPDATE
			SET
				client_id = ?,
				client_secret = ?,
				auth_url = ?,
				auth_url_params = ?,
				token_url = ?,
				redirect_url = ?,
				scopes = ?
		`,
		c.Name,
		c.ClientID, c.ClientSecret, c.AuthURL, authURLParams, c.TokenURL, c.RedirectURL, scopes,
		c.ClientID, c.ClientSecret, c.AuthURL, authURLParams, c.TokenURL, c.RedirectURL, scopes,
	)
	if err != nil {
		return nil, err
	}
	return r.GetConfig(ctx, c.Name)
}

func (r *repo) GetConfig(ctx context.Context, name string) (*oauthmodel.Config, error) {
	row := r.db.QueryRowx(`
		SELECT
			name,
			client_id,
			client_secret,
			auth_url,
			auth_url_params,
			token_url,
			redirect_url,
			scopes,
			created_at,
			updated_at
		FROM config
		WHERE name = ?`,
		name,
	)

	var cr configRow
	err := row.StructScan(&cr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errtype.NotFoundError{Err: err}
		}
		return nil, err
	}

	out, err := configRowToConfig(&cr)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *repo) DeleteConfig(ctx context.Context, name string) error {
	return crudutil.Delete(r.db, `DELETE FROM config WHERE name = ?`, name)
}

func (r *repo) ListProviders(ctx context.Context, args *oauthmodel.ListOAuthProvidersRequest) (*oauthmodel.ListOAuthProvidersResponse, error) {
	sz, err := crudutil.GetPageSize(args.PageSize, 500)
	if err != nil {
		return nil, err
	}

	sb := sq.
		StatementBuilder.
		Select("name, client_id, client_secret, auth_url, auth_url_params, token_url, redirect_url, scopes, created_at, updated_at").
		From("config").
		OrderBy("name ASC").
		Limit(uint64(sz) + 1) // get n+1 so we know if there's a next page

	if args.PageToken != "" {
		page := &listConfigPageData{}
		err := crudutil.DecodePageData(args.PageToken, page)
		if err != nil {
			return nil, err
		}
		sb = sb.Where(sq.GtOrEq{"name": page.NextName})
	}

	sql, sqlArgs, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Queryx(sql, sqlArgs...)
	if err != nil {
		return nil, err
	}

	i := 0
	providers := make([]*oauthmodel.Provider, 0, sz)
	for rows.Next() {
		var r configRow
		err = rows.StructScan(&r)
		if err != nil {
			return nil, err
		}
		c, err := configRowToConfig(&r)
		if err != nil {
			return nil, err
		}
		providers = append(providers, &oauthmodel.Provider{
			Name:   c.Name,
			Config: c,
		})
		i += 1
	}

	nextPageToken := ""
	if len(providers) > int(sz) {
		s, err := crudutil.EncodePageData(&listConfigPageData{
			NextName: providers[len(providers)-1].Name,
		})
		if err != nil {
			return nil, err
		}
		nextPageToken = s
		providers = providers[0 : len(providers)-1]
	}

	return &oauthmodel.ListOAuthProvidersResponse{
		Providers:     providers,
		NextPageToken: nextPageToken,
	}, nil
}

type listConfigPageData struct {
	NextName string `json:"next_name"`
}

func configRowToConfig(r *configRow) (*oauthmodel.Config, error) {
	c, err := time.Parse(sqlite3util.DatetimeFormat, r.CreatedAt)
	if err != nil {
		return nil, err
	}

	u, err := time.Parse(sqlite3util.DatetimeFormat, r.UpdatedAt)
	if err != nil {
		return nil, err
	}

	scopes := []string{}
	err = json.Unmarshal(r.Scopes, &scopes)
	if err != nil {
		return nil, err
	}

	authURLParams := map[string]string{}
	err = json.Unmarshal(r.AuthURLParams, &authURLParams)
	if err != nil {
		return nil, err
	}

	return &oauthmodel.Config{
		Name:          r.Name,
		ClientID:      r.ClientID,
		ClientSecret:  r.ClientSecret,
		AuthURL:       r.AuthURL,
		TokenURL:      r.TokenURL,
		RedirectURL:   r.RedirectURL,
		Scopes:        scopes,
		AuthURLParams: authURLParams,
		CreatedAt:     present.ToAPITime(c),
		UpdatedAt:     present.ToAPITime(u),
	}, nil
}
