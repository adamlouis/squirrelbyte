package model

// resources
type Config struct {
	ClientID                         string      `json:"client_id"`
	ClientSecret                     string      `json:"client_secret"`
	Endpoint                         *Endpoint   `json:"endpoint"`
	RedirectURL                      string      `json:"redirect_url"`
	Scopes                           []string    `json:"scopes"`
	AdditionalAuthorizationURLParams []*URLParam `json:"additional_authorization_url_params"`
}

type URLParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Endpoint struct {
	AuthURL  string `json:"auth_url"`
	TokenURL string `json:"token_url"`
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

// request / response
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
	Providers []*Provider `json:"providers"`
}
