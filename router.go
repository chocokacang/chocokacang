package chocokacang

import (
	"net/http"
)

type Handler func(c *Context)

type Handlers []Handler

type Router struct {
	env *CK
}

func (r *Router) route(method, path string, handlers Handlers) {

}

func (r *Router) Get(path string, handlers ...Handler) {
	r.route(http.MethodGet, path, handlers)
}
