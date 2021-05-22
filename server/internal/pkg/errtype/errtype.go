package errtype

import "net/http"

func Code(err error) int {
	if err == nil {
		return http.StatusInternalServerError
	}
	switch err.(type) {
	case NotFoundError:
		return http.StatusNotFound
	case BadRequestError:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

//go:generate sh -c "./gen.sh NotFoundError BadRequestError > errtype.gen.go"
