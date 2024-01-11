package main

import (
	"net/http"

	"github.com/JuniorPaula/whatsapp-link-generate/api"
	"github.com/go-chi/chi"
)

func routes() http.Handler {
	// http.HandleFunc("/generate-links", api.GenerateLinksHandler)
	mux := chi.NewRouter()

	mux.Post("/generate-links", api.GenerateLinksHandler)

	return mux
}
