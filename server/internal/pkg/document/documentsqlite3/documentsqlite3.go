package documentsqlite3

import (
	"context"
	"errors"
	"fmt"
	"time"

	b64 "encoding/base64"
	"encoding/json"

	sq "github.com/Masterminds/squirrel"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/document"
	jl "github.com/adamlouis/squirrelbyte/server/internal/pkg/document/jsonlogic"
	jls "github.com/adamlouis/squirrelbyte/server/internal/pkg/document/jsonlogic/jsonlogicsqlite3"
	"github.com/jmoiron/sqlx"
)

const (
	// todo: MANY tables, not just `documents` ... user should be able to add aribitrary table
	// todo: add sqlite column index from API ... or just do them all
	// todo: auto-complete document paths
	datetimeFormat = "2006-01-02 15:04:05" // note: no T
	maxPageSize    = uint64(1000)
)

// NewDocumentRepository returns a new document repository
func NewDocumentRepository(db sqlx.Ext) document.Repository {
	return &documentRepo{
		db: db,
	}
}

type documentRepo struct {
	db sqlx.Ext
}

func (dr *documentRepo) Init(ctx context.Context) error {
	// todo: real migration, not this string
	// migrate on startup
	// fine for now
	_, err := dr.db.Exec(`
	CREATE TABLE IF NOT EXISTS document(
		id TEXT NOT NULL UNIQUE CHECK(id <> ''),
		header TEXT NOT NULL CHECK(json_type(header) == 'object'),
		body TEXT NOT NULL CHECK(json_type(body) == 'object'),
		created_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL CHECK(id <> ''),
		updated_at TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL CHECK(id <> '')
	);
	
	
	CREATE TRIGGER IF NOT EXISTS set_document_updated_at
	AFTER UPDATE ON document
	BEGIN
		UPDATE document SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
	END;

	CREATE TABLE IF NOT EXISTS path(
		name TEXT NOT NULL UNIQUE CHECK(name <> '')
	);`)
	return err
}

func (dr *documentRepo) Put(ctx context.Context, d *document.Document) (*document.Document, error) {
	if d.ID == "" {
		return nil, errors.New("id must not be empty")
	}

	_, err := dr.db.Exec(`
			INSERT INTO
				document
					(id, header, body)
				VALUES
					(?, ?, ?)
				ON CONFLICT (id)
				DO UPDATE
					SET header = ?, body = ?`,
		d.ID, d.Header, d.Body, d.Header, d.Body)
	if err != nil {
		return nil, err
	}

	return dr.Get(ctx, d.ID)
}

type documentRow struct {
	ID        string `db:"id"`
	Body      []byte `db:"body"`
	Header    []byte `db:"header"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func (dr *documentRepo) Get(ctx context.Context, documentID string) (*document.Document, error) {
	row := dr.db.QueryRowx(`SELECT id, body, header, created_at, updated_at FROM document WHERE id = ?`, documentID)

	var r documentRow
	err := row.StructScan(&r)
	if err != nil {
		return nil, err
	}

	return docRowToDoc(&r)
}

func (dr *documentRepo) Delete(ctx context.Context, documentID string) error {
	r, err := dr.db.Exec(`DELETE FROM document WHERE id = ?`, documentID)

	if err != nil {
		return err
	}

	ct, err := r.RowsAffected()
	if err != nil {
		return err
	}

	if ct == 0 {
		return errors.New("not found")
	}
	if ct > 1 {
		return errors.New("unexpected")
	}

	return nil
}

func (dr *documentRepo) List(ctx context.Context, args *document.ListDocumentArgs) (*document.ListDocumentResult, error) {
	sz := maxPageSize
	if args.PageSize > 0 {
		sz = uint64(args.PageSize)
	}

	if sz > maxPageSize {
		return nil, errors.New("bad request: page size too large")
	}

	sb := sq.
		StatementBuilder.
		Select("id, body, header, created_at, updated_at").
		From("document").
		OrderBy("id ASC").
		Limit(sz + 1) // get n+1 so we know if there's a next page

	if args.PageToken != "" {
		page := &listDocumentsPageData{}
		err := decodePageData(args.PageToken, page)
		if err != nil {
			return nil, err
		}
		sb = sb.Where(sq.GtOrEq{"id": page.NextID})
	}

	sql, sqlArgs, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := dr.db.Queryx(sql, sqlArgs...)
	if err != nil {
		return nil, err
	}

	documents := make([]*document.Document, 0, sz)

	for rows.Next() {
		var r documentRow
		err = rows.StructScan(&r)
		if err != nil {
			return nil, err
		}
		d, err := docRowToDoc(&r)
		if err != nil {
			return nil, err
		}
		documents = append(documents, d)
	}

	nextPageToken := ""
	if len(documents) > int(sz) {
		s, err := encodePageData(&listDocumentsPageData{
			NextID: documents[len(documents)-1].ID,
		})
		if err != nil {
			return nil, err
		}
		nextPageToken = s
		documents = documents[0 : len(documents)-1]
	}

	return &document.ListDocumentResult{
		Documents: documents,
		PageResult: document.PageResult{
			NextPageToken: nextPageToken,
		},
	}, nil
}

func docRowToDoc(r *documentRow) (*document.Document, error) {
	c, err := time.Parse(datetimeFormat, r.CreatedAt)
	if err != nil {
		return nil, err
	}

	u, err := time.Parse(datetimeFormat, r.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &document.Document{
		ID:        r.ID,
		Header:    r.Header,
		Body:      r.Body,
		CreatedAt: c,
		UpdatedAt: u,
	}, nil
}

// warning: search DOES NOT use prepared statements in order to allow more expressive queries. only use in read-only mode.
// todo: add code-level guard rails to restrict to read-only or other safe contexts
func (dr *documentRepo) Query(ctx context.Context, q *document.Query) (*document.QueryResult, error) {
	sz := maxPageSize
	if q.Limit > 0 {
		sz = uint64(q.Limit)
	}

	if sz > maxPageSize {
		return nil, errors.New("bad request: page size too large")
	}

	sqlizer := jls.NewSQLizer()

	s, err := jl.AllToSQL(sqlizer, q.Select)
	if err != nil {
		return nil, err
	}

	selection := []string{"id", "body", "header", "created_at", "updated_at"}
	// use the json string as the column name in the result
	// it will get overridden if `as` clause is provided
	// no special handling for collisions between names or names & raw
	rawColumnNamesBySelectName := map[string]string{}
	if len(s) > 0 {
		selection = s
		for i := range q.Select {
			b, e := json.Marshal(q.Select[i])
			if e != nil {
				return nil, e
			}
			rawColumnNamesBySelectName[selection[i]] = string(b)
		}
	}

	sb := sq.
		StatementBuilder.
		Select(selection...).
		From("document"). // todo: many arbitrary tables
		Limit(sz + 1)     // get n+1 so we know if there's a next page

	where, err := sqlizer.ToSQL(q.Where)
	if err != nil {
		return nil, err
	}
	sb = sb.Where(where)

	orderBys, err := jl.AllToSQL(sqlizer, q.OrderBy)
	if err != nil {
		return nil, err
	}
	orderBys = append(orderBys, "id ASC") // add id for stable order
	sb = sb.OrderBy(orderBys...)

	if len(q.GroupBy) > 0 {
		groupBys, err := jl.AllToSQL(sqlizer, q.GroupBy)
		if err != nil {
			return nil, err
		}
		sb = sb.GroupBy(groupBys...)
	}

	offset := uint64(0)
	if q.PageToken != "" {
		page := &queryDocumentsPageData{}
		err = decodePageData(q.PageToken, page)
		if err != nil {
			return nil, err
		}
		offset = page.Offset
	}
	sb = sb.Offset(offset)

	sql, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	jsonlog(map[string]interface{}{
		"jsonlogic": q,
		"query":     sql,
		"args":      args,
	})

	rows, err := dr.db.Queryx(sql, args...)
	if err != nil {
		return nil, err
	}

	result := []interface{}{}
	for rows.Next() {
		rowMap := map[string]interface{}{}
		err = rows.MapScan(rowMap)
		if err != nil {
			return nil, err
		}

		resultMap := map[string]interface{}{}
		// try to convert all []byte values to json
		for k, v := range rowMap {
			if kRaw, ok := rawColumnNamesBySelectName[k]; ok {
				k = kRaw
			}
			bs, ok := v.([]byte)
			if ok {
				mp := map[string]interface{}{}
				err := json.Unmarshal(bs, &mp)
				if err != nil {
					return nil, err
				}
				v = mp
			}
			resultMap[k] = v
		}
		result = append(result, resultMap)
	}

	nextPageToken := ""
	if len(result) > int(sz) {
		result = result[0 : len(result)-1]
		s, err := encodePageData(&queryDocumentsPageData{
			Offset: offset + uint64(len(result)),
		})
		if err != nil {
			return nil, err
		}
		nextPageToken = s
	}

	return &document.QueryResult{
		Result: result,
		PageResult: document.PageResult{
			NextPageToken: nextPageToken,
		},
	}, nil
}

type queryDocumentsPageData struct {
	Offset uint64 `json:"offset"`
	// NextID string `json:"next_id"` // TODO: for perf, use id as page cursor if no order by clause is provided
}

type listDocumentsPageData struct {
	NextID string `json:"next_id"`
}

func encodePageData(i interface{}) (string, error) {
	if i == nil {
		return "", nil
	}
	bytes, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return b64.StdEncoding.EncodeToString(bytes), nil
}

func decodePageData(s string, i interface{}) error {
	bytes, err := b64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, i)
}

func jsonlog(i interface{}) {
	b, _ := json.Marshal(i)
	fmt.Println(string(b))
}
