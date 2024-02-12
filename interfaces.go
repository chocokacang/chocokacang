package chocokacang

import "net/http"

type Module interface {
	AddGroupRouter(r func(r *Router))
}

type Response interface {
	http.ResponseWriter
}
