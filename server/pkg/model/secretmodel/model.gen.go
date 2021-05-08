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
	PageToken string `json:"page_token"`
	PageSize  int    `json:"page_size"`
}
type ListSecretsResponse struct {
	Secrets       []*Secret `json:"secrets"`
	NextPageToken string    `json:"next_page_token"`
}
type GetSecretPathParams struct {
	Key string
}
type PutSecretPathParams struct {
	Key string
}
