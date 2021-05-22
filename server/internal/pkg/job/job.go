package job

import (
	"context"

	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

type Repository interface {
	// crud
	Get(ctx context.Context, id string) (*jobmodel.Job, error)
	List(ctx context.Context, args *jobmodel.ListJobsQueryParams) (*jobmodel.ListJobsResponse, error)
	Delete(ctx context.Context, id string) error
	// job queue semantics
	Queue(ctx context.Context, j *jobmodel.Job) (*jobmodel.Job, error)
	Claim(ctx context.Context, opts *ClaimOptions) (*jobmodel.Job, error)
	Release(ctx context.Context, id string) (*jobmodel.Job, error)
	Success(ctx context.Context, id string) (*jobmodel.Job, error)
	Error(ctx context.Context, id string) (*jobmodel.Job, error)
}

type JobStatus string

const (
	JobStatusQueued  JobStatus = "QUEUED"
	JobStatusClaimed JobStatus = "CLAIMED"
	JobStatusSuccess JobStatus = "SUCCESS"
	JobStatusError   JobStatus = "ERROR"
)

type ClaimOptions struct {
	JobID string
	Names []string
}
