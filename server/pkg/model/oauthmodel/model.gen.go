// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: --component model --config ../../config/api.oauth.yml --package oauthmodel --out ./oauthmodel/model.gen.go
package oauthmodel

type Config struct {
	Name          string            `json:"name"`
	ClientID      string            `json:"client_id"`
	ClientSecret  string            `json:"client_secret"`
	AuthURL       string            `json:"auth_url"`
	TokenURL      string            `json:"token_url"`
	RedirectURL   string            `json:"redirect_url"`
	Scopes        []string          `json:"scopes"`
	AuthURLParams map[string]string `json:"auth_url_params"`
	CreatedAt     string            `json:"created_at"`
	UpdatedAt     string            `json:"updated_at"`
}
type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Expiry       string `json:"expiry"`
}
type Provider struct {
	Name   string  `json:"name"`
	Config *Config `json:"config"`
}
type GetOAuthTokenRequest struct {
	Code string `json:"code"`
}
type GetOAuthAuthorizationURLResponse struct {
	URL string `json:"url"`
}
type ListOAuthProvidersRequest struct {
	PageToken string `json:"page_token"`
	PageSize  int    `json:"page_size"`
}
type ListOAuthProvidersResponse struct {
	Providers     []*Provider `json:"providers"`
	NextPageToken string      `json:"next_page_token"`
}
type GetOAuthAuthorizationURLPathParams struct {
	Name string
}
type GetOAuthTokenPathParams struct {
	Name string
}
type GetOAuthConfigPathParams struct {
	Name string
}
type PutOAuthConfigPathParams struct {
	Name string
}
type DeleteOAuthConfigPathParams struct {
	Name string
}
