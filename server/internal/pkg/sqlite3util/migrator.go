package sqlite3util

import (
	"bytes"
	"embed"
	"io"
	"io/fs"
	"sort"
	"strings"
	"time"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/jsonlog"
	"github.com/jmoiron/sqlx"
)

type Migrator interface {
	Up() error
	Down() error
}

func NewMigrator(db sqlx.Ext, fs embed.FS) Migrator {
	return &mgtr{
		db: db,
		fs: fs,
	}
}

type mgtr struct {
	db sqlx.Ext
	fs embed.FS
}

func (m *mgtr) Up() error {
	files, err := getFileNames(m.fs)
	if err != nil {
		return err
	}
	return executeQueryFiles(m.fs, m.db, getWithSuffix(files, ".up.sql"))
}

func (m *mgtr) Down() error {
	files, err := getFileNames(m.fs)
	if err != nil {
		return err
	}
	return executeQueryFiles(m.fs, m.db, getWithSuffix(files, ".down.sql"))
}

func getFileNames(fs embed.FS) ([]string, error) {
	root, err := fs.ReadDir(".")
	if err != nil {
		return nil, err
	}

	files, err := getFileNamesRec(fs, "", root, []string{})
	if err != nil {
		return nil, err
	}

	sort.Strings(files)
	return files, nil
}

// TODO: ugly, tail cursion
func getFileNamesRec(fs embed.FS, path string, entries []fs.DirEntry, results []string) ([]string, error) {
	if len(entries) == 0 {
		return []string{}, nil
	}

	for _, e := range entries {
		entryPath := "'"
		if path != "" {
			entryPath = path + "/" + e.Name()
		} else {
			entryPath = e.Name()
		}

		if e.IsDir() {
			children, err := fs.ReadDir(e.Name())
			if err != nil {
				return nil, err
			}
			r, err := getFileNamesRec(fs, entryPath, children, results)
			if err != nil {
				return nil, err
			}
			results = r
		} else {
			results = append(results, entryPath)
		}
	}

	return results, nil
}

func getWithSuffix(vs []string, sfx string) []string {
	r := []string{}
	for _, v := range vs {
		if strings.HasSuffix(v, sfx) {
			r = append(r, v)
		}
	}
	return r
}

func executeQueryFiles(fs embed.FS, db sqlx.Ext, filenames []string) error {
	for _, filename := range filenames {
		f, err := fs.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()

		buf := bytes.NewBuffer(nil)
		_, err = io.Copy(buf, f)
		if err != nil {
			return err
		}

		jsonlog.Log(
			"name", "Migrate",
			"filename", filename,
			"timestamp", time.Now(),
		)
		_, err = db.Exec(buf.String())
		if err != nil {
			return err
		}
	}
	return nil
}
