package gamehdl

import (
	"encoding/json"
	"hex-structure/internal/core/ports"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HTTPHandler struct {
	gamesService ports.GamesService
}

func NewHTTPHandler(gamesService ports.GamesService) *HTTPHandler {
	return &HTTPHandler{
		gamesService: gamesService,
	}
}

func (hdl *HTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	game, err := hdl.gamesService.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(game); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (hdl *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body BodyCreate

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	game, err := hdl.gamesService.Create(body.Name, body.Size, body.Bombs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(BuildResponseCreate(game)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (hdl *HTTPHandler) RevealCell(w http.ResponseWriter, r *http.Request) {
	var body BodyRevealCell
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := chi.URLParam(r, "id")
	game, err := hdl.gamesService.Reveal(id, body.Row, body.Cell)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(BuildResponseCreate(game)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
