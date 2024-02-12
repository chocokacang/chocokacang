package chocokacang_test

import (
	"testing"

	"github.com/chocokacang/chocokacang"
)

func Module(m chocokacang.Module) {
	m.AddGroupRouter(func(r *chocokacang.Router) {
		r.Get("/", func(c *chocokacang.Context) {})
	})
}

func BenchmarkAddModule(b *testing.B) {

	server := chocokacang.New()

	for i := 0; i < b.N; i++ {
		server.AddModule(Module)
	}
}
