package cgr

import (
	"net/http"
)

type Route struct {
	Path        string
	Methods     []string
	HandlerFunc http.HandlerFunc
	Router      *Router
	Middleware  []*Middleware
}

func (r *Route) Method(methods ...string) *Route {
	r.Methods = methods
	return r
}

func (r *Route) Handler(handler http.HandlerFunc) *Route {
	r.HandlerFunc = handler
	return r
}

func (r *Route) Insert() {
	r.Router.Routes.Insert(r)
}
