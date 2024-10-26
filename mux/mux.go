package mux

import (
	"fmt"
	"net/http"

	"github.com/panda-mod/web/middleware"
)

// Route instance
type Route struct {
	Pattern string
	Method  string
	Handler http.Handler
}

// Router interface
type Router interface {
	Group(pattern string) Router
	Use(middlewares ...middleware.Middleware) Router
	GET(pattern string, fn http.HandlerFunc) Router
	HEAD(pattern string, fn http.HandlerFunc) Router
	POST(pattern string, fn http.HandlerFunc) Router
	PUT(pattern string, fn http.HandlerFunc) Router
	DELETE(pattern string, fn http.HandlerFunc) Router
	PATCH(pattern string, fn http.HandlerFunc) Router
	OPTIONS(pattern string, fn http.HandlerFunc) Router
	Any(pattern string, fn http.HandlerFunc) Router
	Handle(pattern string, fn http.Handler, methods ...string) Router
	HandleFunc(pattern string, fn http.HandlerFunc, methods ...string) Router
}

// defaultMethodNotAllowed 405 Method Not Allowed
func defaultMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, http.StatusText(http.StatusMethodNotAllowed))
}
