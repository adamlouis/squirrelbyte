package present

import (
	"encoding/json"
	"time"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/document"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/job"
	"github.com/adamlouis/squirrelbyte/server/pkg/model"
)

// why this indirection here
// seems like overkill with just one resource
// yagni?

// APIDocumentToInternalDocument returns the API representation of document form the internal representation
func APIDocumentToInternalDocument(d *model.Document) (*document.Document, error) {
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
func InternalDocumentToAPIDocument(d *document.Document) (*model.Document, error) {
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

	return &model.Document{
		ID:        d.ID,
		Header:    h,
		Body:      b,
		CreatedAt: ToAPITime(d.CreatedAt),
		UpdatedAt: ToAPITime(d.UpdatedAt),
	}, nil
}

// InternalDocumentsToAPIDocuments returns the internal representation of documents from the API representation
func InternalDocumentsToAPIDocuments(ds []*document.Document) ([]*model.Document, error) {
	if ds == nil {
		return nil, nil
	}

	r := make([]*model.Document, len(ds))
	for i := range ds {
		a, err := InternalDocumentToAPIDocument(ds[i])
		if err != nil {
			return nil, err
		}
		r[i] = a
	}
	return r, nil
}

// ToAPITime returns an RFC3339 time from a golangtim
func ToAPITime(t time.Time) string {
	return t.Format(time.RFC3339)
}

// APIJobToInternalJob returns the API representation of job form the internal representation
func APIJobToInternalJob(j *model.Job) (*job.Job, error) {

	input, err := json.Marshal(j.Input)
	if err != nil {
		return nil, err
	}

	output, _ := json.Marshal(j.Output)
	if err != nil {
		return nil, err
	}

	return &job.Job{
		ID:     j.ID,
		Name:   j.Name,
		Status: "", // todo
		Input:  input,
		Output: &output,
	}, nil
}

// InternalJobToAPIJob returns the internal representation of a job from the API representation
func InternalJobToAPIJob(j *job.Job) (*model.Job, error) {

	var input map[string]interface{}
	err := json.Unmarshal(j.Input, &input)
	if err != nil {
		return nil, err
	}

	var output map[string]interface{}
	if j.Output != nil {
		err = json.Unmarshal(*j.Output, &output)
		if err != nil {
			return nil, err
		}
	}

	var s, e *string
	if j.SucceededAt != nil {
		s = sptr(ToAPITime(*j.SucceededAt))
	}
	if j.ErroredAt != nil {
		e = sptr(ToAPITime(*j.ErroredAt))
	}

	return &model.Job{
		ID:          j.ID,
		Name:        j.Name,
		Status:      string(j.Status),
		Input:       input,
		Output:      output,
		SucceededAt: s,
		ErroredAt:   e,
		CreatedAt:   ToAPITime(j.CreatedAt),
		UpdatedAt:   ToAPITime(j.UpdatedAt),
	}, nil
}

// InternalJobsToAPIJobs returns the internal representation of jobs from the API representation
func InternalJobsToAPIJobs(js []*job.Job) ([]*model.Job, error) {
	if js == nil {
		return nil, nil
	}
	r := make([]*model.Job, len(js))
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
