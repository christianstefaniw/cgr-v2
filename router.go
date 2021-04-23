package cgr

import (
	"context"
	"fmt"
	"net/http"
)

type Router struct {
	Routes *Tree
}

type Key int

const ParamsKey = Key(iota)

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path

	result, _ := r.Routes.Search(path, method)

	if result.Params != nil {
		ctx := context.WithValue(req.Context(), ParamsKey, result.Params)
		req = req.WithContext(ctx)
	}

	result.Route.HandlerFunc.ServeHTTP(w, req)
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
