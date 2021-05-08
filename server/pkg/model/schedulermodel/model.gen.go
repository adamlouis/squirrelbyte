// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: '--component model --config ../../config/api.scheduler.yml --package schedulermodel --out ./schedulermodel/model.gen.go'

package schedulermodel

type JSONObject map[string]interface{}
type Scheduler struct {
	JobName  string     `json:"job_name"`
	Input    JSONObject `json:"input"`
	Schedule string     `json:"schedule"`
}
type ListSchedulersRequest struct {
	PageSize  int    `json:"page_size"`
	PageToken string `json:"page_token"`
}
type ListSchedulersResponse struct {
	NextPageToken string       `json:"next_page_token"`
	Schedulers    []*Scheduler `json:"schedulers"`
}
type DeleteSchedulerPathParams struct {
	SchedulerID string
}
type GetSchedulerPathParams struct {
	SchedulerID string
}
type PutSchedulerPathParams struct {
	SchedulerID string
}
