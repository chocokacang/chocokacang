package chocokacang

import (
	"net/http"

	"github.com/chocokacang/chocokacang/log"
)

var _ Response = (*writer)(nil)

type writer struct {
	http.ResponseWriter
	size   int
	status int
}

func (w *writer) init(rsw http.ResponseWriter) {
	w.ResponseWriter = rsw
}

func (w *writer) WriteHeader(code int) {
	if code > 0 && w.status != code {
		log.Debug(log.WARNING, "Headers were already written. Wanted to override status code %d with %d", w.status, code)
	}

	w.status = code
}
