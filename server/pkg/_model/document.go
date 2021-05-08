package model

// Document is the document resource
type Document struct {
	ID        string     `json:"id"`
	Header    JSONObject `json:"header"`
	Body      JSONObject `json:"body"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
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
	DocumentID string
}

// PutDocumentPathParams are the path parameters for putting documents
type PutDocumentPathParams struct {
	DocumentID string
}

// DeleteDocumentPathParams are that path params for deleting documents
type DeleteDocumentPathParams struct {
	DocumentID string
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
	Result        []interface{} `json:"result"`
	NextPageToken string        `json:"next_page_token"`
	Insights      JSONObject    `json:"insights"`
}
