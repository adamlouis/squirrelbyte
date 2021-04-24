package main

// // Status is the status response
// type Status struct {
// 	Status    string `json:"status"`
// 	Timestamp string `json:"timestamp"`
// }

// // Error is the standard error response
// type Error struct {
// 	Message string `json:"message"`
// }

// // Document is the document resource
// type Document struct {
// 	ID        string                 `json:"id"`
// 	Header    map[string]interface{} `json:"header"`
// 	Body      map[string]interface{} `json:"body"`
// 	CreatedAt string                 `json:"created_at"`
// 	UpdatedAt string                 `json:"updated_at"`
// }

// // Job is the job resource
// type Job struct {
// 	ID     string                 `json:"id"`
// 	Name   string                 `json:"name"`
// 	Status string                 `json:"status"`
// 	Input  map[string]interface{} `json:"input"`
// 	Output map[string]interface{} `json:"output"`

// 	SucceededAt *string `json:"succeeded_at"`
// 	ErroredAt   *string `json:"errored_at"`
// 	CreatedAt   string  `json:"created_at"`
// 	UpdatedAt   string  `json:"updated_at"`
// }

// // ListDocumentsQueryParams are the query params for listing documents
// type ListDocumentsQueryParams struct {
// 	PageToken string `json:"page_token"`
// 	PageSize  int    `json:"page_size"`
// }

// // ListDocumentsResponse is the response for listing documents
// type ListDocumentsResponse struct {
// 	Documents     []*Document `json:"documents"`
// 	NextPageToken string      `json:"next_page_token"`
// }

// // GetJobPathParams are the path params for getting jobs
// type GetJobPathParams struct {
// 	JobID string
// }

// // GetDocumentPathParams are the path params for getting documents
// type GetDocumentPathParams struct {
// 	DocumentID string
// }

// // PutDocumentPathParams are the path parameters for putting documents
// type PutDocumentPathParams struct {
// 	DocumentID string
// }

// // DeleteDocumentPathParams are that path params for deleting documents
// type DeleteDocumentPathParams struct {
// 	DocumentID string
// }

// // QueryDocumentsRequest is the request body for querying documents
// type QueryDocumentsRequest struct {
// 	Select    []interface{} `json:"select"`
// 	GroupBy   []interface{} `json:"group_by"`
// 	OrderBy   []interface{} `json:"order_by"`
// 	Where     interface{}   `json:"where"`
// 	Limit     int           `json:"limit"`
// 	PageToken string        `json:"page_token"`
// }

// // QueryDocumentsResponse is the response body from querying documents
// type QueryDocumentsResponse struct {
// 	Result        []interface{}          `json:"result"`
// 	NextPageToken string                 `json:"next_page_token"`
// 	Insights      map[string]interface{} `json:"insights"`
// }

// type ClaimJobRequest struct {
// 	Names []string
// }
// type ClaimJobPathParams struct {
// 	JobID string
// }
// type DeleteJobPathParams struct {
// 	JobID string
// }
// type ReleaseJobPathParams struct {
// 	JobID string
// }
// type SetJobSuccessPathParams struct {
// 	JobID string
// }
// type SetJobErrorPathParams struct {
// 	JobID string
// }

// // ListJobsQueryParams are the query params for listing documents
// type ListJobsQueryParams struct {
// 	PageToken string `json:"page_token"`
// 	PageSize  int    `json:"page_size"`
// }

// // ListJobsResponse is the response for listing documents
// type ListJobsResponse struct {
// 	Jobs          []*Job `json:"jobs"`
// 	NextPageToken string `json:"next_page_token"`
// }
