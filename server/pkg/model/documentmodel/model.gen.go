// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: '--component model --config ../../config/api.document.yml --package documentmodel --out ./documentmodel/model.gen.go'

package documentmodel

type JSONObject map[string]interface{}
type Document struct {
	Body      JSONObject `json:"body"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
	ID        string     `json:"id"`
	Header    JSONObject `json:"header"`
}
type ListDocumentsQueryParams struct {
	PageToken string `json:"page_token"`
	PageSize  int    `json:"page_size"`
}
type ListDocumentsResponse struct {
	Documents     []*Document `json:"documents"`
	NextPageToken string      `json:"next_page_token"`
}
type QueryDocumentsRequest struct {
	Select    []interface{} `json:"select"`
	GroupBy   []interface{} `json:"group_by"`
	OrderBy   []interface{} `json:"order_by"`
	Where     interface{}   `json:"where"`
	Limit     int           `json:"limit"`
	PageToken string        `json:"page_token"`
}
type QueryDocumentsResponse struct {
	Result        []interface{} `json:"result"`
	NextPageToken string        `json:"next_page_token"`
	Insights      JSONObject    `json:"insights"`
}
type GetDocumentPathParams struct {
	DocumentID string
}
type PutDocumentPathParams struct {
	DocumentID string
}
type DeleteDocumentPathParams struct {
	DocumentID string
}
