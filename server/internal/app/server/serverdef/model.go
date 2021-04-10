package serverdef

// TODO: generate all of `serverdef` package from conf / openapi declaration

// resources
type Status struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
type Error struct {
	Message string `json:"message"`
}

type Document struct {
	ID        string                 `json:"id"`
	Header    map[string]interface{} `json:"header"`
	Body      map[string]interface{} `json:"body"`
	CreatedAt string                 `json:"created_at"`
	UpdatedAt string                 `json:"updated_at"`
}

// list documents
type ListDocumentsQueryParams struct {
	PageToken string `json:"page_token"`
	PageSize  int    `json:"page_size"`
}
type ListDocumentsResponse struct {
	Documents     []*Document `json:"documents"`
	NextPageToken string      `json:"next_page_token"`
}

// get document
type GetDocumentPathParams struct {
	DocumentID string `json:"document_id"`
}

// put document
type PutDocumentPathParams struct {
	DocumentID string `json:"document_id"`
}

// delete document
type DeleteDocumentPathParams struct {
	DocumentID string `json:"document_id"`
}

// query documents
type SearchDocumentsRequest struct {
	Select    []interface{} `json:"select"`
	GroupBy   []interface{} `json:"group_by"`
	OrderBy   []interface{} `json:"order_by"`
	Where     interface{}   `json:"where"`
	Limit     int           `json:"limit"`
	PageToken string        `json:"page_token"`
}

type SearchDocumentsResponse struct {
	Result        []interface{}          `json:"result"`
	NextPageToken string                 `json:"next_page_token"`
	Insights      map[string]interface{} `json:"insights"`
}
