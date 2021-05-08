// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: '--component model --config ../../config/api.scheduler.yml --package schedulermodel --out ./schedulermodel/model.gen.go'

package schedulermodel

type JSONObject map[string]interface{}
type Scheduler struct {
	Input    JSONObject `json:"input"`
	Schedule string     `json:"schedule"`
	JobName  string     `json:"job_name"`
}
type ListSchedulersRequest struct {
	PageToken string `json:"page_token"`
	PageSize  int    `json:"page_size"`
}
type ListSchedulersResponse struct {
	Schedulers    []*Scheduler `json:"schedulers"`
	NextPageToken string       `json:"next_page_token"`
}
type PutSchedulerPathParams struct {
	SchedulerID string
}
type DeleteSchedulerPathParams struct {
	SchedulerID string
}
type GetSchedulerPathParams struct {
	SchedulerID string
}
