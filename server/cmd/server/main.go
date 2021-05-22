package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/adamlouis/squirrelbyte/server/internal/app/server/documentserver"
	"github.com/adamlouis/squirrelbyte/server/internal/app/server/jobserver"
	"github.com/adamlouis/squirrelbyte/server/internal/app/server/kvserver"
	"github.com/adamlouis/squirrelbyte/server/internal/app/server/oauthserver"
	"github.com/adamlouis/squirrelbyte/server/internal/app/server/schedulerserver"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/document/documentsqlite3"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/errtype"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/job/jobsqlite3"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/jsonlog"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/kv/kvsqlite3"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/oauth/oauthsqlite3"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/scheduler/schedulersqlite3"
	"github.com/adamlouis/squirrelbyte/server/internal/pkg/sqlite3util"
	"github.com/adamlouis/squirrelbyte/server/pkg/client/jobclient"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

// TODO: revisit config-based init
const (
	_envServerPort             = "SQUIRRELBYTE_SERVER_PORT"
	_envSQLiteConnectionString = "SQUIRRELBYTE_SQLITE3_CONNECTION_STRING"
	_envStaticDir              = "SQUIRRELBYTE_STATIC_DIR"
	_envAllowedHTTPMethods     = "SQUIRRELBYTE_ALLOWED_HTTP_METHODS"
	_envAllowedHTTPPaths       = "SQUIRRELBYTE_ALLOWED_HTTP_PATHS"

	defaultServerPort = 9922
)

type config struct {
	ServerPort         int
	SQLite3Path        string
	StaticDir          string
	AllowedHTTPMethods map[string]bool
	AllowedHTTPPaths   map[string]bool
}

func newConfig() (*config, error) {
	if fDotenv != nil && *fDotenv != "" {
		err := godotenv.Load(*fDotenv)
		if err != nil {
			return nil, err
		}
	}

	serverPort := defaultServerPort
	pstr := os.Getenv(_envServerPort)
	if pstr != "" {
		p, err := strconv.Atoi(pstr)
		if err != nil {
			return nil, err
		}
		serverPort = p
	}

	methods := strings.Split(os.Getenv(_envAllowedHTTPMethods), ",")
	allowedMethods := map[string]bool{}
	for _, m := range methods {
		allowedMethods[m] = true
	}
	paths := strings.Split(os.Getenv(_envAllowedHTTPPaths), ",")
	allowedPaths := map[string]bool{}
	for _, p := range paths {
		allowedPaths[p] = true
	}

	return &config{
		ServerPort:         serverPort,
		SQLite3Path:        os.Getenv(_envSQLiteConnectionString),
		StaticDir:          os.Getenv(_envStaticDir),
		AllowedHTTPMethods: allowedMethods,
		AllowedHTTPPaths:   allowedPaths,
	}, nil
}

func newDB(c *config, path string) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1) // TODO: use RW lock or WAL rather than 1 max conn
	return db, nil
}

var (
	fDotenv = flag.String("dotenv", "", "a .env file from which to read environment variables. useful for local development.")
)

func main() {
	ctx := context.Background()

	flag.Parse()

	c, err := newConfig()
	if err != nil {
		log.Fatal(err)
	}

	documentDB, err := newDB(c, "./data/document.db")
	if err != nil {
		log.Fatal(err)
	}
	defer documentDB.Close()

	jobDB, err := newDB(c, "./data/job.db")
	if err != nil {
		log.Fatal(err)
	}
	defer jobDB.Close()

	oauthDB, err := newDB(c, "./data/oauth.db")
	if err != nil {
		log.Fatal(err)
	}
	defer oauthDB.Close()

	kvDB, err := newDB(c, "./data/kv.db")
	if err != nil {
		log.Fatal(err)
	}
	defer kvDB.Close()

	schedulerDB, err := newDB(c, "./data/scheduler.db")
	if err != nil {
		log.Fatal(err)
	}
	defer schedulerDB.Close()

	if err := sqlite3util.NewMigrator(documentDB, documentsqlite3.MigrationFS).Up(); err != nil {
		log.Fatal(err)
	}
	if err := sqlite3util.NewMigrator(jobDB, jobsqlite3.MigrationFS).Up(); err != nil {
		log.Fatal(err)
	}
	if err := sqlite3util.NewMigrator(oauthDB, oauthsqlite3.MigrationFS).Up(); err != nil {
		log.Fatal(err)
	}
	if err := sqlite3util.NewMigrator(kvDB, kvsqlite3.MigrationFS).Up(); err != nil {
		log.Fatal(err)
	}
	if err := sqlite3util.NewMigrator(schedulerDB, schedulersqlite3.MigrationFS).Up(); err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	documentserver.RegisterRouter(documentserver.NewAPIHandler(documentDB), apiRouter, errtype.Code)
	jobserver.RegisterRouter(jobserver.NewAPIHandler(jobDB), apiRouter, errtype.Code)
	oauthserver.RegisterRouter(oauthserver.NewAPIHandler(oauthDB), apiRouter, errtype.Code)
	kvserver.RegisterRouter(kvserver.NewAPIHandler(kvDB), apiRouter, errtype.Code)
	schedulerserver.RegisterRouter(
		schedulerserver.NewAPIHandler(
			ctx,
			schedulerDB,
			jobclient.NewHTTPClient("http://localhost:9922/api"),
		),
		apiRouter,
		errtype.Code,
	)
	// TODO: revisit auth
	router.Use(loggerMiddleware)

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%d", c.ServerPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r)
		// TODO: do w/ tracing / opentelemetry, start a span, pass down context, etc
		// TODO: produce just 1 structured event per req w/ all metadata
		jsonlog.Log(
			"name", fmt.Sprintf("%s:%s", r.Method, r.URL.Path),
			"type", "REQUEST",
			"method", r.Method,
			"duration_ms", time.Since(now)/time.Millisecond,
			"path", r.URL.Path,
			"time", time.Now().Format(time.RFC3339),
		)
	})
}
