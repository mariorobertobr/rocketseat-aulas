package api

import (
	"encoding/json"
	"net/http"
	omd "project3/omdb"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(apikey string) http.Handler {

	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/movies", handleGetMovies(apikey))

	return r
}
func handleGetMovies(apikey string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("s")
		result, err := omd.SearchResult(apikey, search)
		if err != nil {

			sendJson(w, Response{Error: err.Error()}, http.StatusInternalServerError)
			return
		}
		sendJson(w, Response{Data: result}, http.StatusOK)
	}
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func sendJson(w http.ResponseWriter, resp any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}
