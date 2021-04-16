package present

import (
	"encoding/json"
	"time"

	"github.com/adamlouis/squirrelbyte/server/internal/app/server/serverdef"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/document"
)

// why this indirection here
// seems like overkill with just one resource
// yagni?

// APIDocumentToInternalDocument returns the API representation of document form the internal representation
func APIDocumentToInternalDocument(d *serverdef.Document) (*document.Document, error) {
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
func InternalDocumentToAPIDocument(d *document.Document) (*serverdef.Document, error) {
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

	return &serverdef.Document{
		ID:        d.ID,
		Header:    h,
		Body:      b,
		CreatedAt: ToAPITime(d.CreatedAt),
		UpdatedAt: ToAPITime(d.UpdatedAt),
	}, nil
}

// InternalDocumentsToAPIDocuments returns the internal representation of documents from the API representation
func InternalDocumentsToAPIDocuments(ds []*document.Document) ([]*serverdef.Document, error) {
	if ds == nil {
		return nil, nil
	}

	r := make([]*serverdef.Document, len(ds))
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
