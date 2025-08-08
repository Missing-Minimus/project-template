package routes

import "net/http"

type Middleware func(http.Handler) http.Handler

func HandlerChain(h http.HandlerFunc, middlewares ...Middleware) http.Handler {
	var handler http.Handler = h
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
