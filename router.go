package cgr

import (
	"net/http"
)

type router struct {
	routes *tree
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path

	route := r.routes.search(path, method)

	route.handler.ServeHTTP(w, req)
}

func (r *router) Route(path, method string, handler http.HandlerFunc) {
	newRoute := route{path: path, method: method, handler: handler}
	r.routes.insert(newRoute)
}

func (r *router) Run(port string) {
	http.ListenAndServe(":"+port, r)
}

func NewRouter() *router {
	tree := newTree()
	return &router{routes: tree}
}
