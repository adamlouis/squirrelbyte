package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/adamlouis/squirrelbyte/server/internal/app/server"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

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

func newDB(c *config) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", c.SQLite3Path)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1) // TODO: use RW lock rather than max conns

	return db, nil
}

var (
	fDotenv = flag.String("dotenv", "", "a .env file from which to read environment variables. useful for local development.")
)

func main() {
	flag.Parse()

	c, err := newConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := newDB(c)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // nolint

	fmt.Printf("starting server :%d\n", c.ServerPort)
	err = server.New().Serve(&server.Opts{
		Port:      c.ServerPort,
		DB:        db,
		StaticDir: c.StaticDir,
		Middlewares: []mux.MiddlewareFunc{
			func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					now := time.Now()
					next.ServeHTTP(w, r)
					// todo: produce just 1 structured event per req w/ all metadata
					// todo: do w/ tracing / opentelemetry, start a span, pass down context, etc
					j, _ := json.Marshal(map[string]interface{}{
						"type":        "REQUEST",
						"method":      r.Method,
						"name":        fmt.Sprintf("%s:%s", r.URL.Path, r.Method),
						"duration_ms": time.Since(now) / time.Millisecond,
						"path":        r.URL.Path,
						"time":        time.Now().Format(time.RFC3339),
					})
					fmt.Println(string(j))
				})
			},
			func(next http.Handler) http.Handler {
				// allow any req with "allowed methods" OR "allowed paths"
				// used to block reqs in read only mode
				// one day, proper authz
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if !c.AllowedHTTPMethods[r.Method] && !c.AllowedHTTPPaths[r.URL.Path] {
						w.Header().Add("Content-Type", "application/json")
						w.WriteHeader(http.StatusForbidden)
						_, _ = w.Write([]byte(`{"message":"forbidden"}`))
						return
					}
					next.ServeHTTP(w, r)
				})
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}
