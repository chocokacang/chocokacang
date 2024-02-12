package chocokacang

import "net/http"

type Context struct {
	Request  *http.Request
	Response Response
	env      *CK
	writer   writer
}

func (c *Context) init() {
	c.Response = &c.writer

}
