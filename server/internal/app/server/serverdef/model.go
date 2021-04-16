package serverdef

// TODO: generate all of `serverdef` package from conf / openapi declaration

// Status is the status response
type Status struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

// Error is the standard error response
type Error struct {
	Message string `json:"message"`
}

// Document is the document resource
type Document struct {
	ID        string                 `json:"id"`
	Header    map[string]interface{} `json:"header"`
	Body      map[string]interface{} `json:"body"`
	CreatedAt string                 `json:"created_at"`
	UpdatedAt string                 `json:"updated_at"`
}

// ListDocumentsQueryParams are the query params for listing documents
type ListDocumentsQueryParams struct {
	PageToken string `json:"page_token"`
	PageSize  int    `json:"page_size"`
}

// ListDocumentsResponse is the response for listing documents
type ListDocumentsResponse struct {
	Documents     []*Document `json:"documents"`
	NextPageToken string      `json:"next_page_token"`
}

// GetDocumentPathParams are the path params for getting documents
type GetDocumentPathParams struct {
	DocumentID string `json:"document_id"`
}

// PutDocumentPathParams are the path parameters for putting documents
type PutDocumentPathParams struct {
	DocumentID string `json:"document_id"`
}

// DeleteDocumentPathParams are that path params for deleting documents
type DeleteDocumentPathParams struct {
	DocumentID string `json:"document_id"`
}

// QueryDocumentsRequest is the request body for querying documents
type QueryDocumentsRequest struct {
	Select    []interface{} `json:"select"`
	GroupBy   []interface{} `json:"group_by"`
	OrderBy   []interface{} `json:"order_by"`
	Where     interface{}   `json:"where"`
	Limit     int           `json:"limit"`
	PageToken string        `json:"page_token"`
}

// QueryDocumentsResponse is the response body from querying documents
type QueryDocumentsResponse struct {
	Result        []interface{}          `json:"result"`
	NextPageToken string                 `json:"next_page_token"`
	Insights      map[string]interface{} `json:"insights"`
}
