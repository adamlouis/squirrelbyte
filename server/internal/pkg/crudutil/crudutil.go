package crudutil

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Delete(db sqlx.Ext, q string, args ...interface{}) error {
	result, err := db.Exec(q, args...)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("not found")
	}
	if count > 1 {
		return errors.New("unexpected: found more than one row matching criteria")
	}
	return nil
}

func EncodePageData(i interface{}) (string, error) {
	if i == nil {
		return "", nil
	}
	bytes, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

func DecodePageData(s string, i interface{}) error {
	bytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, i)
}

func GetPageSize(requestedSz, maxSz int) (int64, error) {
	if requestedSz > 0 {
		if requestedSz > maxSz {
			return 0, fmt.Errorf("max page size is %d", maxSz)
		}
		return int64(requestedSz), nil
	}
	return int64(maxSz), nil
}
