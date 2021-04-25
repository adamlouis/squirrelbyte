package model

// Job is the job resource
type Job struct {
	ID     string                 `json:"id"`
	Name   string                 `json:"name"`
	Status string                 `json:"status"`
	Input  map[string]interface{} `json:"input"`
	Output map[string]interface{} `json:"output"`

	SucceededAt *string `json:"succeeded_at"`
	ErroredAt   *string `json:"errored_at"`
	ClaimedAt   *string `json:"claimed_at"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// GetJobPathParams are the path params for getting jobs
type GetJobPathParams struct {
	JobID string
}

type ClaimJobRequest struct {
	Names []string
}

type ClaimJobPathParams struct {
	JobID string
}

type DeleteJobPathParams struct {
	JobID string
}

type ReleaseJobPathParams struct {
	JobID string
}

type SetJobSuccessPathParams struct {
	JobID string
}

type SetJobErrorPathParams struct {
	JobID string
}

// ListJobsQueryParams are the query params for listing documents
type ListJobsQueryParams struct {
	PageToken string `json:"page_token"`
	PageSize  int    `json:"page_size"`
}

// ListJobsResponse is the response for listing documents
type ListJobsResponse struct {
	Jobs          []*Job `json:"jobs"`
	NextPageToken string `json:"next_page_token"`
}
