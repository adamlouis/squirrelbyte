package serverdef

import (
	"encoding/json"
	"net/http"
)

// TODO: generate all of `serverdef` package from conf / openapi declaration

// SendError sends an error response
func SendError(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")

	code := http.StatusInternalServerError
	if httpErr, ok := err.(*HTTPError); ok {
		code = int(httpErr.Code())
	}

	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(&Error{
		Message: err.Error(),
	})
}

// SendOK sends an success response
func SendOK(w http.ResponseWriter, body interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(body)
}
