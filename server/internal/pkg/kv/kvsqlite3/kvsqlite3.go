package kvsqlite3

import (
	"context"
	"database/sql"
	"embed"
	"encoding/json"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/crudutil"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/errtype"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/kv"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/present"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/sqlite3util"
	"github.com/adamlouis/squirrelbyte/server/pkg/model/kvmodel"
	"github.com/jmoiron/sqlx"
)

func NewRepository(db sqlx.Ext) kv.Repository {
	return &repo{db: db}
}

type repo struct {
	db sqlx.Ext
}

//go:embed migration/*.sql
var MigrationFS embed.FS

type kvRow struct {
	Key       string `db:"key"`
	Value     []byte `db:"value"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func (r *repo) Put(ctx context.Context, in *kvmodel.KV) (*kvmodel.KV, error) {
	rowStruct, err := kvToKVRow(in)
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(`
			INSERT INTO
				kv
					(key, value)
				VALUES
					(?, ?)
				ON CONFLICT (key)
				DO UPDATE SET
				value = ?`,
		rowStruct.Key, rowStruct.Value,
		rowStruct.Value,
	)
	if err != nil {
		return nil, err
	}

	return r.Get(ctx, in.Key)
}

func (r *repo) Get(ctx context.Context, key string) (*kvmodel.KV, error) {
	row := r.db.QueryRowx(`SELECT key, value, created_at, updated_at FROM kv WHERE key = ?`, key)
	var rowStruct kvRow
	err := row.StructScan(&rowStruct)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errtype.NotFoundError{Err: err}
		}
		return nil, err
	}
	return kvRowToKV(&rowStruct)
}

func (r *repo) List(ctx context.Context, args *kvmodel.ListKVsRequest) (*kvmodel.ListKVsResponse, error) {
	sz, err := crudutil.GetPageSize(args.PageSize, 500)
	if err != nil {
		return nil, err
	}

	sb := sq.
		StatementBuilder.
		Select("key, value, created_at, updated_at").
		From("kv").
		OrderBy("key ASC").
		Limit(uint64(sz) + 1) // get n+1 so we know if there's a next page

	if args.PageToken != "" {
		page := &listKVsPageData{}
		err := crudutil.DecodePageData(args.PageToken, page)
		if err != nil {
			return nil, err
		}
		sb = sb.Where(sq.GtOrEq{"key": page.NextKey})
	}

	sql, sqlArgs, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Queryx(sql, sqlArgs...)
	if err != nil {
		return nil, err
	}

	kvs := make([]*kvmodel.KV, 0, sz)

	for rows.Next() {
		var r kvRow
		err = rows.StructScan(&r)
		if err != nil {
			return nil, err
		}
		kv, err := kvRowToKV(&r)
		if err != nil {
			return nil, err
		}
		kvs = append(kvs, kv)
	}

	nextPageToken := ""
	if len(kvs) > int(sz) {
		s, err := crudutil.EncodePageData(&listKVsPageData{
			NextKey: kvs[len(kvs)-1].Key,
		})
		if err != nil {
			return nil, err
		}
		nextPageToken = s
		kvs = kvs[0 : len(kvs)-1]
	}

	return &kvmodel.ListKVsResponse{
		Kvs:           kvs,
		NextPageToken: nextPageToken,
	}, nil
}

type listKVsPageData struct {
	NextKey string `json:"next_key"`
}

func (r *repo) Delete(ctx context.Context, key string) error {
	return crudutil.Delete(r.db, "DELETE FROM kv WHERE key = ?", key)
}

func kvRowToKV(row *kvRow) (*kvmodel.KV, error) {
	c, err := time.Parse(sqlite3util.DatetimeFormat, row.CreatedAt)
	if err != nil {
		return nil, err
	}

	u, err := time.Parse(sqlite3util.DatetimeFormat, row.UpdatedAt)
	if err != nil {
		return nil, err
	}

	var value interface{}
	err = json.Unmarshal(row.Value, &value)
	if err != nil {
		return nil, err
	}

	return &kvmodel.KV{
		Key:       row.Key,
		Value:     value,
		CreatedAt: present.ToAPITime(c),
		UpdatedAt: present.ToAPITime(u),
	}, nil
}

func kvToKVRow(kv *kvmodel.KV) (*kvRow, error) {
	if kv.Key == "" {
		return nil, errors.New("key must not be empty")
	}

	value, err := json.Marshal(kv.Value)
	if err != nil {
		return nil, err
	}
	if len(value) == 0 {
		return nil, errors.New("value must not be empty")
	}

	return &kvRow{
		Key:   kv.Key,
		Value: value,
	}, nil
}
