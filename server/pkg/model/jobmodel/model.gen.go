// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: '--component model --config ../../config/api.job.yml --package jobmodel --out ./jobmodel/model.gen.go'

package jobmodel

type JSONObject map[string]interface{}
type ClaimSomeJobRequest struct {
	Names []string `json:"names"`
}
type Job struct {
	UpdatedAt    string     `json:"updated_at"`
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Status       string     `json:"status"`
	Input        JSONObject `json:"input"`
	ClaimedAt    *string    `json:"claimed_at"`
	CreatedAt    string     `json:"created_at"`
	ScheduledFor *string    `json:"scheduled_for"`
	SucceededAt  *string    `json:"succeeded_at"`
	ErroredAt    *string    `json:"errored_at"`
}
type ListJobsQueryParams struct {
	PageSize  int    `json:"page_size"`
	PageToken string `json:"page_token"`
}
type ListJobsResponse struct {
	Jobs          []*Job `json:"jobs"`
	NextPageToken string `json:"next_page_token"`
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
