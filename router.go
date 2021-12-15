package cgr

import (
	"context"
	"fmt"
	"net/http"
)

type Router struct {
	Routes *Tree
}

type key int

// Key for url parameters in context
const ParamsKey = key(iota)

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer internalError(w)

	method := req.Method
	path := req.URL.Path

	result, err := r.Routes.Search(path, method)

	if err != nil {
		http.Error(w, fmt.Sprintf("Access %s: %s", path, err), http.StatusNotImplemented)
		return
	}

	if result.Params != nil {
		ctx := context.WithValue(req.Context(), ParamsKey, result.Params)
		req = req.WithContext(ctx)
	}

	result.Route.ExecuteRoute(w, req)
}

func internalError(w http.ResponseWriter) {
	if r := recover(); r != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}

func (r *Router) Route(path string) *Route {
	return &Route{Path: path, Router: r}
}

func (r *Router) Run(port string) {
	fmt.Println("Listening on:", port)
	http.ListenAndServe(":"+port, r)
}

func NewRouter() *Router {
	tree := NewTree()
	return &Router{Routes: tree}
}
