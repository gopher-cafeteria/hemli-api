package router

import (
	v1 "hemli/internal/router/v1"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	Router *chi.Mux
}

func NewRouter(r *chi.Mux) *Router {
	return &Router{
		Router: r,
	}
}

func (r *Router) MountHandlers() {
	// Mount all Middleware here
	r.Router.Use(middleware.Logger)

	r.Router.Route("/v1", v1.Handlers)
}
