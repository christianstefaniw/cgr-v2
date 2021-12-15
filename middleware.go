package cgr

import "net/http"

type MiddlewareHandler http.HandlerFunc

type Middleware struct {
	handler MiddlewareHandler
}

// Create middleware for passed in handler
func NewMiddleware(handler MiddlewareHandler) *Middleware {
	return &Middleware{
		handler: handler,
	}
}

// Run the current middleware
func (m *Middleware) Run(w http.ResponseWriter, r *http.Request) {
	m.handler(w, r)
}
