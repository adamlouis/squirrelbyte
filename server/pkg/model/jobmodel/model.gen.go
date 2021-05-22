// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: --component model --config ../../config/api.job.yml --package jobmodel --out ./jobmodel/model.gen.go
package jobmodel

type JSONObject map[string]interface{}
type Job struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Status       string     `json:"status"`
	Input        JSONObject `json:"input"`
	ScheduledFor *string    `json:"scheduled_for"`
	SucceededAt  *string    `json:"succeeded_at"`
	ErroredAt    *string    `json:"errored_at"`
	ClaimedAt    *string    `json:"claimed_at"`
	CreatedAt    string     `json:"created_at"`
	UpdatedAt    string     `json:"updated_at"`
}
type ListJobsQueryParams struct {
	PageSize  int    `json:"page_size"`
	PageToken string `json:"page_token"`
}
type ListJobsResponse struct {
	Jobs          []*Job `json:"jobs"`
	NextPageToken string `json:"next_page_token"`
}
type ClaimSomeJobRequest struct {
	Names []string `json:"names"`
}
type GetJobPathParams struct {
	JobID string
}
type DeleteJobPathParams struct {
	JobID string
}
type ClaimJobPathParams struct {
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
