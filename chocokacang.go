package chocokacang

import (
	"net/http"
	"sync"

	"github.com/chocokacang/chocokacang/config"
	"github.com/chocokacang/chocokacang/log"
)

type CK struct {
	Router *Router
	trees  trees
	pool   sync.Pool
}

func New() *CK {
	// Create environment instance
	env := &CK{}

	// Create router instance
	env.Router = &Router{
		env: env,
	}

	// Create pool for context
	env.pool.New = func() any {
		return &Context{env: env}
	}

	return env
}

func (env *CK) AddRouter() {

}

func (env *CK) AddGroupRouter(load func(r *Router)) {
	load(env.Router)
}

func (env *CK) AddModule(load func(m Module)) {
	load(env)
}

func (env *CK) ServeHTTP(rsw http.ResponseWriter, rq *http.Request) {
	// Get context instance from the pool
	c := env.pool.Get().(*Context)

	rsw.WriteHeader(http.StatusNotFound)
	rsw.WriteHeader(http.StatusOK)

	rsw.Write([]byte("X"))

	// Put context to the pool
	env.pool.Put(c)
}

func (env *CK) Run() {
	port := config.Get.AppPort
	log.Info("Run the server in localhost:%s", port)

	server := &http.Server{
		Addr:     ":" + port,
		Handler:  env,
		ErrorLog: log.Logger(log.WARNING),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Error("Stop the server, %v", err)
	}
}
