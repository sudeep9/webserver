package webserver

import "github.com/go-chi/chi/v5"

type Handlers interface {
	Register(path string, r *chi.Mux)
}
