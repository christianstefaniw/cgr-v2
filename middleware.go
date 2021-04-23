package cgr

import "net/http"

type Middleware struct {
	handler http.HandlerFunc
}
