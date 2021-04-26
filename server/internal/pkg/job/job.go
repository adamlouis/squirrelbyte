package job

import (
	"context"
	"time"
)

type Repository interface {
	// crud
	Init(ctx context.Context) error
	Get(ctx context.Context, id string) (*Job, error)
	List(ctx context.Context, args *ListJobArgs) (*ListJobResults, error)
	Delete(ctx context.Context, id string) error
	// job queue semantics
	Queue(ctx context.Context, j *Job) (*Job, error)
	Claim(ctx context.Context, opts ClaimOptions) (*Job, error)
	Release(ctx context.Context, id string) (*Job, error)
	Success(ctx context.Context, id string) (*Job, error)
	Error(ctx context.Context, id string) (*Job, error)
}

type JobStatus string

const (
	JobStatusQueued  JobStatus = "QUEUED"
	JobStatusClaimed JobStatus = "CLAIMED"
	JobStatusSuccess JobStatus = "SUCCESS"
	JobStatusError   JobStatus = "ERROR"
)

type Job struct {
	ID     string
	Name   string
	Status JobStatus

	Input []byte

	ScheduledFor *time.Time
	SucceededAt  *time.Time
	ErroredAt    *time.Time
	ClaimedAt    *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ClaimOptions struct {
	JobID string
	Names []string
}

// PageArgs are the arguments getting a page
type PageArgs struct {
	PageSize  int
	PageToken string
}

// PageResult are the values getting the next page
type PageResult struct {
	NextPageToken string
}

type ListJobArgs struct {
	PageArgs
}

type ListJobResults struct {
	PageResult
	Jobs []*Job
}
