package routes

import "net/http"

type Middleware func(http.Handler) http.Handler

func HandlerChain(h http.HandlerFunc, middlewares ...Middleware) http.Handler {
	var handler http.Handler = h                 // isso aqui releva, apenas type check
	for i := len(middlewares) - 1; i >= 0; i-- { // aqui eu pego a lista de middlewares e coloco o controller dentro deles
		handler = middlewares[i](handler)
	}
	return handler // aqui eu retorno todos os middlewares + a minha rota
}
