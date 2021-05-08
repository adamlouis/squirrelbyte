// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: '--component model --config ../../config/api.secret.yml --package secretmodel --out ./secretmodel/model.gen.go'

package secretmodel

type Secret struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}
type ListSecretsRequest struct {
	PageSize  int    `json:"page_size"`
	PageToken string `json:"page_token"`
}
type ListSecretsResponse struct {
	NextPageToken string    `json:"next_page_token"`
	Secrets       []*Secret `json:"secrets"`
}
type GetSecretPathParams struct {
	Key string
}
type PutSecretPathParams struct {
	Key string
}
