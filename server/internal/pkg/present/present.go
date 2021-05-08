package present

import (
	"encoding/json"
	"time"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/document"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/job"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel"
)

// why this indirection here
// seems like overkill with just one resource
// yagni?

// APIDocumentToInternalDocument returns the API representation of document form the internal representation
func APIDocumentToInternalDocument(d *documentmodel.Document) (*document.Document, error) {
	if d == nil {
		return nil, nil
	}

	h, err := json.Marshal(d.Header)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(d.Body)
	if err != nil {
		return nil, err
	}

	return &document.Document{
		ID:     d.ID,
		Header: h,
		Body:   b,
	}, nil
}

// InternalDocumentToAPIDocument returns the internal representation of a document from the API representation
func InternalDocumentToAPIDocument(d *document.Document) (*documentmodel.Document, error) {
	if d == nil {
		return nil, nil
	}

	var h map[string]interface{}
	err := json.Unmarshal(d.Header, &h)
	if err != nil {
		return nil, err
	}

	var b map[string]interface{}
	err = json.Unmarshal(d.Body, &b)
	if err != nil {
		return nil, err
	}

	return &documentmodel.Document{
		ID:        d.ID,
		Header:    h,
		Body:      b,
		CreatedAt: ToAPITime(d.CreatedAt),
		UpdatedAt: ToAPITime(d.UpdatedAt),
	}, nil
}

// InternalDocumentsToAPIDocuments returns the internal representation of documents from the API representation
func InternalDocumentsToAPIDocuments(ds []*document.Document) ([]*documentmodel.Document, error) {
	if ds == nil {
		return nil, nil
	}

	r := make([]*documentmodel.Document, len(ds))
	for i := range ds {
		a, err := InternalDocumentToAPIDocument(ds[i])
		if err != nil {
			return nil, err
		}
		r[i] = a
	}
	return r, nil
}

// ToAPITime returns an RFC3339 time from a golang time
func ToAPITime(t time.Time) string {
	return t.Format(time.RFC3339)
}

// ToInternalTime returns a golang time from a RFC3339
func ToInternalTime(s string) (time.Time, error) {
	// fmt.Println(time.RFC3339[:len(time.RFC3339)-5], s)
	// return time.Parse(time.RFC3339[:len(time.RFC3339)-5], s)
	return time.Parse(time.RFC3339, s)
}

// APIJobToInternalJob returns the API representation of job form the internal representation
func APIJobToInternalJob(j *jobmodel.Job) (*job.Job, error) {

	input, err := json.Marshal(j.Input)
	if err != nil {
		return nil, err
	}

	var scheduledFor *time.Time
	if j.ScheduledFor != nil {
		sf, err := ToInternalTime(*j.ScheduledFor)
		if err != nil {
			return nil, err
		}
		scheduledFor = &sf
	}

	return &job.Job{
		ID:           j.ID,
		Name:         j.Name,
		Status:       "", // todo
		Input:        input,
		ScheduledFor: scheduledFor,
	}, nil
}

// InternalJobToAPIJob returns the internal representation of a job from the API representation
func InternalJobToAPIJob(j *job.Job) (*jobmodel.Job, error) {

	var input map[string]interface{}
	err := json.Unmarshal(j.Input, &input)
	if err != nil {
		return nil, err
	}

	var s, e, c, sf *string
	if j.SucceededAt != nil {
		s = sptr(ToAPITime(*j.SucceededAt))
	}
	if j.ErroredAt != nil {
		e = sptr(ToAPITime(*j.ErroredAt))
	}
	if j.ClaimedAt != nil {
		c = sptr(ToAPITime(*j.ClaimedAt))
	}
	if j.ScheduledFor != nil {
		sf = sptr(ToAPITime(*j.ScheduledFor))
	}

	return &jobmodel.Job{
		ID:           j.ID,
		Name:         j.Name,
		Status:       string(j.Status),
		Input:        input,
		SucceededAt:  s,
		ErroredAt:    e,
		ClaimedAt:    c,
		ScheduledFor: sf,
		CreatedAt:    ToAPITime(j.CreatedAt),
		UpdatedAt:    ToAPITime(j.UpdatedAt),
	}, nil
}

// InternalJobsToAPIJobs returns the internal representation of jobs from the API representation
func InternalJobsToAPIJobs(js []*job.Job) ([]*jobmodel.Job, error) {
	if js == nil {
		return nil, nil
	}
	r := make([]*jobmodel.Job, len(js))
	for i := range js {
		j, err := InternalJobToAPIJob(js[i])
		if err != nil {
			return nil, err
		}
		r[i] = j
	}
	return r, nil
}

func sptr(s string) *string {
	return &s
}
