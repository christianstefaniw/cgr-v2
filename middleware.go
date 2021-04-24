package cgr

import "net/http"

type MiddlewareHandler http.HandlerFunc

type Middleware struct {
	handler MiddlewareHandler
}

func NewMiddleware(handler MiddlewareHandler) *Middleware {
	return &Middleware{
		handler: handler,
	}
}

func (m *Middleware) Run(w http.ResponseWriter, r *http.Request) {
	m.handler(w, r)
}
