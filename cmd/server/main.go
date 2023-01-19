package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/ducminhgd/gohook/pkg/handlers"
)

func main() {
	r := chi.NewRouter()

	// Uses middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/gitlab", func(r chi.Router) {
		r.Post("/{hookID}", handlers.HandleGitlabHook)
	})

	http.ListenAndServe(":8000", r)
}
