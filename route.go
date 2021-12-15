package cgr

import (
	"net/http"
)

// Structure for a Route
type Route struct {
	Path string

	// Methods allowed on this route
	Methods []string

	// Route handler
	HandlerFunc http.HandlerFunc

	// Handler that this route belongs to
	Router *Router

	// Middleware that should be run when this route is requested
	Middleware []*Middleware

	// Stop request from continuing if it is a preflight check.
	//
	// For example, a preflight request that does not have the correct body and,
	// headers might set off a handler which will cause an error
	Preflight bool
}

// Add valid HTTP methods for a certian route
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

// Insert route into the router
func (r *Route) Insert() {
	r.Router.Routes.Insert(r)
}

// assign middleware to route
func (r *Route) Assign(middleware ...*Middleware) *Route {
	r.Middleware = append(r.Middleware, middleware...)
	return r
}

func (r *Route) ExecuteRoute(w http.ResponseWriter, req *http.Request) {

	// run all middleware
	for _, m := range r.Middleware {
		m.Run(w, req)
	}

	// check if it is a preflight request and if user wants
	// to halt execution if it is a preflight request
	if req.Method == "OPTIONS" && r.Preflight {

		// run preflight handler
		CorsHandler(w, req)

		/// halt
		return
	}

	// server route
	r.HandlerFunc.ServeHTTP(w, req)
}
