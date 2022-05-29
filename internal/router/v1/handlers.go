package v1

import "github.com/go-chi/chi/v5"

func Handlers(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", ListUsers)
		r.Post("/", CreateUser)
	})
}
