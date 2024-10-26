package middleware

import "net/http"

// Middleware 中间件
type Middleware func(next http.Handler) http.Handler

// Chain 构造一个中间件链，并返回一个 http.Handler
func Chain(middlewares []Middleware, h http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
