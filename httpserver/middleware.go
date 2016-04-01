package httpserver

import "net/http"

//go:generate counterfeiter . Middleware
type Middleware interface {
	Wrap(http.Handler) http.Handler
}

//go:generate counterfeiter . MiddlewareWrapper
type MiddlewareWrapper interface {
	AddMiddlewareToHandler(http.Handler) http.Handler
}
