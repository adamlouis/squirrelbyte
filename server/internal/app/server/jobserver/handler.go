package jobserver

import "github.com/jmoiron/sqlx"

func NewAPIHandler(db *sqlx.DB) APIHandler {
	return &hdl{
		db: db,
	}
}

type hdl struct {
	db *sqlx.DB
}
