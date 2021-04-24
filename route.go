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
	Preflight   bool
}

func (r *Route) Method(methods ...string) *Route {
	r.Methods = methods
	return r
}

func (r *Route) HandlePreflight() *Route {
	r.Preflight = true
	return r
}

func (r *Route) Handler(handler http.HandlerFunc) *Route {
	r.HandlerFunc = handler
	return r
}

func (r *Route) Insert() {
	r.Router.Routes.Insert(r)
}

func (r *Route) Assign(middleware ...*Middleware) *Route {
	r.Middleware = append(r.Middleware, middleware...)
	return r
}

func (r *Route) ExecuteRoute(w http.ResponseWriter, req *http.Request) {
	for _, m := range r.Middleware {
		m.Run(w, req)
	}

	if req.Method == "OPTIONS" && r.Preflight {
		CorsHandler(w, req)
		return
	}

	r.HandlerFunc.ServeHTTP(w, req)
}
