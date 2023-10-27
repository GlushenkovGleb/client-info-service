package httpcontroller

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(ch *ClientController) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/clients", func(r chi.Router) {
		r.Get("/", ch.GetClients)
		r.Post("/", ch.AddClient)
		r.Delete("/{clientId}/", ch.DeleteClient)
		r.Patch("/{clientId}/", ch.UpdateClient)
	})
	return r
}
