// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: '--component model --config ../../config/api.oauth.yml --package oauthmodel --out ./oauthmodel/model.gen.go'

package oauthmodel

type Config struct {
	ClientID                         string      `json:"client_id"`
	ClientSecret                     string      `json:"client_secret"`
	Endpoint                         *Endpoint   `json:"endpoint"`
	RedirectURL                      string      `json:"redirect_url"`
	Scopes                           []string    `json:"scopes"`
	AdditionalAuthorizationURLParams []*URLParam `json:"additional_authorization_url_params"`
}
type Token struct {
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Expiry       string `json:"expiry"`
	AccessToken  string `json:"access_token"`
}
type GetOAuthTokenRequest struct {
	Code string `json:"code"`
}
type GetOAuthAuthorizationURLResponse struct {
	URL string `json:"url"`
}
type ListOAuthProvidersResponse struct {
	NextPageToken string      `json:"next_page_token"`
	Providers     []*Provider `json:"providers"`
}
type URLParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type Endpoint struct {
	AuthURL  string `json:"auth_url"`
	TokenURL string `json:"token_url"`
}
type Provider struct {
	Name   string  `json:"name"`
	Config *Config `json:"config"`
}
type ListOAuthProvidersRequest struct {
	PageToken string `json:"page_token"`
	PageSize  int    `json:"page_size"`
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
