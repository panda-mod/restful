package mux

import (
	"net/http"
	"path"
	"strings"

	"github.com/panda-mod/kit/slice"
	"github.com/panda-mod/web/middleware"
)

type router struct {
	prefix      string
	engine      *Engine
	middlewares []middleware.Middleware
}

func (r *router) Group(pattern string) Router {
	return &router{
		prefix:      path.Join(r.prefix, pattern),
		engine:      r.engine,
		middlewares: r.middlewares,
	}
}

func (r *router) Use(middlewares ...middleware.Middleware) Router {
	r.middlewares = append(r.middlewares, middlewares...)
	return r
}

func (r *router) GET(pattern string, fn http.HandlerFunc) Router {
	return r.Handle(pattern, fn, http.MethodGet)
}

func (r *router) HEAD(pattern string, fn http.HandlerFunc) Router {
	return r.Handle(pattern, fn, http.MethodHead)
}

func (r *router) POST(pattern string, fn http.HandlerFunc) Router {
	return r.Handle(pattern, fn, http.MethodPost)
}

func (r *router) PUT(pattern string, fn http.HandlerFunc) Router {
	return r.Handle(pattern, fn, http.MethodPut)
}

func (r *router) DELETE(pattern string, fn http.HandlerFunc) Router {
	return r.Handle(pattern, fn, http.MethodDelete)
}

func (r *router) PATCH(pattern string, fn http.HandlerFunc) Router {
	return r.Handle(pattern, fn, http.MethodPatch)
}

func (r *router) OPTIONS(pattern string, fn http.HandlerFunc) Router {
	return r.Handle(pattern, fn, http.MethodOptions)
}

func (r *router) Any(pattern string, fn http.HandlerFunc) Router {
	return r.Handle(pattern, fn)
}

func (r *router) HandleFunc(pattern string, fn http.HandlerFunc, methods ...string) Router {
	return r.Handle(pattern, fn, methods...)
}

// Handle 注册路由
func (r *router) Handle(pattern string, handler http.Handler, methods ...string) Router {
	pattern = "/" + strings.TrimLeft(r.prefix+pattern, "/")
	r.addRoute(pattern, handler, methods...)
	r.addHandle(pattern, handler, methods...)
	return r
}

// addRoute 记录路由信息
func (r *router) addRoute(pattern string, handler http.Handler, methods ...string) {
	r.engine.routes = append(r.engine.routes, &Route{
		Pattern: pattern,
		Handler: handler,
		Method:  slice.First(methods, "*"),
	})
}

// addHandle 添加到ServeMux
func (r *router) addHandle(pattern string, handler http.Handler, methods ...string) {
	if pattern == "/" {
		pattern = pattern + "{$}"
	}
	pattern = strings.Join(append(methods, pattern), " ")
	r.engine.serveMux.Handle(pattern, middleware.Chain(r.middlewares, handler))
}
