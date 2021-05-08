// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: '--component model --config ../../config/api.job.yml --package jobmodel --out ./jobmodel/model.gen.go'

package jobmodel

type JSONObject map[string]interface{}
type Job struct {
	ScheduledFor *string    `json:"scheduled_for"`
	SucceededAt  *string    `json:"succeeded_at"`
	ClaimedAt    *string    `json:"claimed_at"`
	CreatedAt    string     `json:"created_at"`
	UpdatedAt    string     `json:"updated_at"`
	ID           string     `json:"id"`
	Status       string     `json:"status"`
	Input        JSONObject `json:"input"`
	Name         string     `json:"name"`
	ErroredAt    *string    `json:"errored_at"`
}
type ListJobsQueryParams struct {
	PageSize  int    `json:"page_size"`
	PageToken string `json:"page_token"`
}
type ListJobsResponse struct {
	NextPageToken string `json:"next_page_token"`
	Jobs          []*Job `json:"jobs"`
}
type ClaimSomeJobRequest struct {
	Names []string `json:"names"`
}
type DeleteJobPathParams struct {
	JobID string
}
type GetJobPathParams struct {
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
