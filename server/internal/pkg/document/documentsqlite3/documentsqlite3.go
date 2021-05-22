package documentsqlite3

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"time"

	"encoding/json"

	sq "github.com/Masterminds/squirrel"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/crudutil"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/document"
	jl "github.com/adamlouis/squirrelbyte/server/internal/pkg/document/jsonlogic"
	jls "github.com/adamlouis/squirrelbyte/server/internal/pkg/document/jsonlogic/jsonlogicsqlite3"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/errtype"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/jsonlog"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/sqlite3util"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel"
	"github.com/jmoiron/sqlx"
)

// todo: MANY tables, not just `documents` ... user should be able to add aribitrary table
// todo: add sqlite column index from API ... or just do them all
// todo: auto-complete document paths

//go:embed migration/*.sql
var MigrationFS embed.FS

// NewDocumentRepository returns a new document repository
func NewDocumentRepository(db sqlx.Ext) document.Repository {
	return &documentRepo{
		db: db,
	}
}

type documentRepo struct {
	db sqlx.Ext
}

type documentRow struct {
	ID        string `db:"id"`
	Body      []byte `db:"body"`
	Header    []byte `db:"header"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func (dr *documentRepo) Put(ctx context.Context, d *documentmodel.Document) (*documentmodel.Document, error) {
	if d.ID == "" {
		return nil, errors.New("id must not be empty")
	}

	h, err := json.Marshal(d.Header)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(d.Body)
	if err != nil {
		return nil, err
	}

	_, err = dr.db.Exec(`
			INSERT INTO
				document
					(id, header, body)
				VALUES
					(?, ?, ?)
				ON CONFLICT (id)
				DO UPDATE
					SET header = ?, body = ?`,
		d.ID,
		h, b,
		h, b,
	)
	if err != nil {
		return nil, err
	}

	return dr.Get(ctx, d.ID)
}

func (dr *documentRepo) Get(ctx context.Context, documentID string) (*documentmodel.Document, error) {
	row := dr.db.QueryRowx(`SELECT id, body, header, created_at, updated_at FROM document WHERE id = ?`, documentID)

	var r documentRow
	err := row.StructScan(&r)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errtype.NotFoundError{Err: err}
		}
		return nil, err
	}

	return docRowToDoc(&r)
}

func (dr *documentRepo) Delete(ctx context.Context, documentID string) error {
	return crudutil.Delete(dr.db, `DELETE FROM document WHERE id = ?`, documentID)
}

func (dr *documentRepo) List(ctx context.Context, args *documentmodel.ListDocumentsQueryParams) (*documentmodel.ListDocumentsResponse, error) {
	sz, err := crudutil.GetPageSize(args.PageSize, 1000)
	if err != nil {
		return nil, err
	}

	sb := sq.
		StatementBuilder.
		Select("id, body, header, created_at, updated_at").
		From("document").
		OrderBy("id ASC").
		Limit(uint64(sz) + 1) // get n+1 so we know if there's a next page

	if args.PageToken != "" {
		page := &listDocumentsPageData{}
		err := crudutil.DecodePageData(args.PageToken, page)
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

	documents := make([]*documentmodel.Document, 0, sz)

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
		s, err := crudutil.EncodePageData(&listDocumentsPageData{
			NextID: documents[len(documents)-1].ID,
		})
		if err != nil {
			return nil, err
		}
		nextPageToken = s
		documents = documents[0 : len(documents)-1]
	}

	return &documentmodel.ListDocumentsResponse{
		Documents:     documents,
		NextPageToken: nextPageToken,
	}, nil
}

func docRowToDoc(r *documentRow) (*documentmodel.Document, error) {
	c, err := time.Parse(sqlite3util.DatetimeFormat, r.CreatedAt)
	if err != nil {
		return nil, err
	}

	u, err := time.Parse(sqlite3util.DatetimeFormat, r.UpdatedAt)
	if err != nil {
		return nil, err
	}

	var h map[string]interface{}
	err = json.Unmarshal(r.Header, &h)
	if err != nil {
		return nil, err
	}

	var b map[string]interface{}
	err = json.Unmarshal(r.Body, &b)
	if err != nil {
		return nil, err
	}

	return &documentmodel.Document{
		ID:        r.ID,
		Header:    h,
		Body:      b,
		CreatedAt: present.ToAPITime(c),
		UpdatedAt: present.ToAPITime(u),
	}, nil
}

// warning: search DOES NOT use prepared statements in order to allow more expressive queries. only use in read-only mode.
// todo: add code-level guard rails to restrict to read-only or other safe contexts
func (dr *documentRepo) Query(ctx context.Context, q *documentmodel.QueryDocumentsRequest) (*documentmodel.QueryDocumentsResponse, error) {
	sz, err := crudutil.GetPageSize(q.Limit, 1000)
	if err != nil {
		return nil, err
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
		From("document").     // todo: many arbitrary tables
		Limit(uint64(sz) + 1) // get n+1 so we know if there's a next page

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
		err = crudutil.DecodePageData(q.PageToken, page)
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

	jsonlog.Log(
		"name", "DocumentQuery",
		"jsonlogic", q,
		"query", sql,
		"args", args,
		"timestamp", time.Now(),
	)

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
		s, err := crudutil.EncodePageData(&queryDocumentsPageData{
			Offset: offset + uint64(len(result)),
		})
		if err != nil {
			return nil, err
		}
		nextPageToken = s
	}

	return &documentmodel.QueryDocumentsResponse{
		Result:        result,
		NextPageToken: nextPageToken,
	}, nil
}

type queryDocumentsPageData struct {
	Offset uint64 `json:"offset"`
	// NextID string `json:"next_id"` // TODO: for perf, use id as page cursor if no order by clause is provided
}

type listDocumentsPageData struct {
	NextID string `json:"next_id"`
}
