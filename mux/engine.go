package mux

import (
	"bytes"
	"net/http"

	"github.com/panda-mod/web/middleware"
)

type Engine struct {
	Router
	routes           []*Route
	serveMux         *http.ServeMux
	NotFound         http.HandlerFunc
	MethodNotAllowed http.HandlerFunc
}

func New() *Engine {
	e := &Engine{
		routes:           make([]*Route, 0),
		serveMux:         http.NewServeMux(),
		NotFound:         http.NotFound,
		MethodNotAllowed: defaultMethodNotAllowed,
	}
	e.Router = &router{
		prefix:      "/",
		engine:      e,
		middlewares: make([]middleware.Middleware, 0),
	}
	return e
}

// Routes returns all registered routes
func (e *Engine) Routes() []*Route {
	return e.routes
}

// ServeHTTP handles HTTP requests
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	writer := &ResponseWriter{
		writer:     w,
		buffer:     &bytes.Buffer{},
		statusCode: http.StatusOK,
	}
	e.serveMux.ServeHTTP(writer, r)
	switch writer.statusCode {
	case http.StatusNotFound:
		writer.buffer.Reset()
		e.NotFound(writer, r)
	case http.StatusMethodNotAllowed:
		writer.buffer.Reset()
		e.MethodNotAllowed(writer, r)
	}
	writer.Finally()
}
