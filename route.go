package cgr

import "net/http"

type route struct {
	method, path string
	handler      http.HandlerFunc
	middleware   []*Middleware
}
