package serverdef

import (
	"net/http"

	"github.com/gorilla/mux"
)

// TODO: generate all of `serverdef` package from conf / openapi declaration

// RegisterRouter registers the router with the api handler
func RegisterRouter(a APIHandler, r *mux.Router) {
	h := apiHandlerToHTTPAPIHandler(a)
	// TODO: split up documents / jobs / secrets / etc
	r.Handle("/status", http.HandlerFunc(h.GetStatus)).Methods(http.MethodGet)
	// documents
	r.Handle("/documents", http.HandlerFunc(h.ListDocuments)).Methods(http.MethodGet)
	r.Handle("/documents", http.HandlerFunc(h.CreateDocument)).Methods(http.MethodPost)
	r.Handle("/documents/{documentID}", http.HandlerFunc(h.GetDocument)).Methods(http.MethodGet)
	r.Handle("/documents/{documentID}", http.HandlerFunc(h.PutDocument)).Methods(http.MethodPut)
	r.Handle("/documents/{documentID}", http.HandlerFunc(h.DeleteDocument)).Methods(http.MethodDelete)
	r.Handle("/documents:query", http.HandlerFunc(h.QueryDocuments)).Methods(http.MethodPost)
	// jobs
	r.Handle("/jobs", http.HandlerFunc(h.ListJobs)).Methods(http.MethodGet)
	r.Handle("/jobs/{jobID}", http.HandlerFunc(h.GetJob)).Methods(http.MethodGet)
	r.Handle("/jobs/{jobID}", http.HandlerFunc(h.DeleteJob)).Methods(http.MethodDelete)
	r.Handle("/jobs:queue", http.HandlerFunc(h.QueueJob)).Methods(http.MethodPost)
	r.Handle("/jobs:claim", http.HandlerFunc(h.ClaimSomeJob)).Methods(http.MethodPost)
	r.Handle("/jobs/{jobID}:claim", http.HandlerFunc(h.ClaimJob)).Methods(http.MethodPost)
	r.Handle("/jobs/{jobID}:release", http.HandlerFunc(h.ReleaseJob)).Methods(http.MethodPost)
	r.Handle("/jobs/{jobID}:success", http.HandlerFunc(h.SetJobSuccess)).Methods(http.MethodPost)
	r.Handle("/jobs/{jobID}:error", http.HandlerFunc(h.SetJobError)).Methods(http.MethodPost)
	// scheduler
	r.Handle("/schedulers", http.HandlerFunc(h.GetStatus)).Methods(http.MethodGet)                  // TODO
	r.Handle("/schedulers", http.HandlerFunc(h.GetStatus)).Methods(http.MethodPost)                 // TODO
	r.Handle("/schedulers/{schedulerID}", http.HandlerFunc(h.GetStatus)).Methods(http.MethodGet)    // TODO
	r.Handle("/schedulers/{schedulerID}", http.HandlerFunc(h.GetStatus)).Methods(http.MethodPut)    // TODO
	r.Handle("/schedulers/{schedulerID}", http.HandlerFunc(h.GetStatus)).Methods(http.MethodDelete) // TODO
	// secrets
	r.Handle("/secrets", http.HandlerFunc(h.ListSecrets)).Methods(http.MethodGet)     // TODO
	r.Handle("/secrets/{key}", http.HandlerFunc(h.GetSecret)).Methods(http.MethodPut) // TODO
	r.Handle("/secrets/{key}", http.HandlerFunc(h.SetSecret)).Methods(http.MethodGet) // TODO
	// oauth
	r.Handle("/oauth/providers", http.HandlerFunc(h.ListOAuthProviders)).Methods(http.MethodGet)                        // TODO
	r.Handle("/oauth/providers/{name}/authorize", http.HandlerFunc(h.GetOAuthAuthorizationURL)).Methods(http.MethodGet) // TODO
	r.Handle("/oauth/providers/{name}/token", http.HandlerFunc(h.GetOAuthToken)).Methods(http.MethodPost)               // TODO
	r.Handle("/oauth/providers/{name}/config", http.HandlerFunc(h.GetOAuthConfig)).Methods(http.MethodGet)              // TODO
	r.Handle("/oauth/providers/{name}/config", http.HandlerFunc(h.SetOAuthConfig)).Methods(http.MethodPost)             // TODO
}
