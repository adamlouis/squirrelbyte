package model

type Secret struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type ListSecretsRequest struct {
	PageToken string `json:"page_token"`
	PageSize  int    `json:"page_size"`
}

type ListSecretsResponse struct {
	Providers []*Provider `json:"providers"`
}
