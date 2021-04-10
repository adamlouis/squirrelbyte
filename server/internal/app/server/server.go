package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/adamlouis/squirrelbyte/server/internal/app/server/serverdef"
	"github.com/adamlouis/squirrelbyte/server/internal/app/server/serverhandler"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// Server is a squirrelbyte API server
type Server interface {
	Serve(o *Opts) error
}

// MiddlewareFn is a middleware function
type MiddlewareFn = mux.MiddlewareFunc

// Opts are options for running the server
type Opts struct {
	Port        int
	DB          *sqlx.DB
	Middlewares []MiddlewareFn
	StaticDir   string
}

// New returns a new server
func New() Server {
	return &srv{}
}

type srv struct{}

func (s *srv) Serve(opts *Opts) error {
	if opts == nil {
		return fmt.Errorf("`opts` may not be nil")
	}

	handler, err := serverhandler.NewAPIHandler(opts.DB)
	if err != nil {
		return err
	}

	// top level router
	router := mux.NewRouter().UseEncodedPath()
	router.Use(opts.Middlewares...)

	// /api router
	apiRouter := router.PathPrefix("/api").Subrouter()
	serverdef.RegisterRouter(handler, apiRouter)

	// static file router
	spa := spaHandler{staticPath: opts.StaticDir, indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%d", opts.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}

// ----------------------------------------------------------------------------
// COPIED FROM github.com/gorilla/mux DOCS: https://github.com/gorilla/mux#serving-single-page-applications
// ----------------------------------------------------------------------------
// spaHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The path to the static directory and
// path to the index file within that static directory are used to
// serve the SPA in the given static directory.
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
