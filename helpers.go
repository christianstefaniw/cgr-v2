package cgr

import (
	"errors"
	"net/http"
)

func deleteEmpty(s []string) []string {
	var removed []string

	for _, str := range s {
		if str != "" {
			removed = append(removed, str)
		}
	}

	return removed
}

// Get parameter from request
func GetParam(r *http.Request, key string) string {
	params, _ := r.Context().Value(ParamsKey).(Params)

	for _, param := range params {
		if param.Key == key {
			return param.Value
		}
	}
	return ""
}

// Literally just an empty handler helper
func EmptyHandler(w http.ResponseWriter, r *http.Request) {}

// Respond with success if request was a preflight request
func CorsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func HandlerNotRegisted() (*SearchResult, error) {
	return nil, errors.New("handler is not registered")
}
