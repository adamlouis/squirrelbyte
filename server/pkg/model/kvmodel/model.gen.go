// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: --component model --config ../../config/api.kv.yml --package kvmodel --out ./kvmodel/model.gen.go
package kvmodel

type KV struct {
	Key       string      `json:"key"`
	Value     interface{} `json:"value"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
}
type ListKVsRequest struct {
	PageToken string `json:"page_token"`
	PageSize  int    `json:"page_size"`
}
type ListKVsResponse struct {
	Kvs           []*KV  `json:"kvs"`
	NextPageToken string `json:"next_page_token"`
}
type GetKVPathParams struct {
	Key string
}
type PutKVPathParams struct {
	Key string
}
type DeleteKVPathParams struct {
	Key string
}
