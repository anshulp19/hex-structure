package main

import (
	"hex-structure/internal/core/services/gameserv"
	"hex-structure/internal/handlers/gamehdl"
	"hex-structure/internal/repositories/gamesrepo"
	"hex-structure/pkg/uidgen"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	gamesRepositories := gamesrepo.NewMemKVS()
	gamesService := gameserv.New(gamesRepositories, uidgen.New())
	gamesHandler := gamehdl.NewHTTPHandler(gamesService)

	r := chi.NewRouter()
	r.Get("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("beep"))
	})
	r.Get("/games/{id}", gamesHandler.Get)
	r.Post("/games", gamesHandler.Create)
	r.Put("/games/{id}", gamesHandler.RevealCell)

	http.ListenAndServe(":3000", r)
}
