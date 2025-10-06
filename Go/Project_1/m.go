package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type User struct {
	Username string
	Id       int64 `json:",string"`
	Role     string
	Password string `json:"-"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func sendJson(w http.ResponseWriter, resp Response, status int) {
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(data)
}

func main() {

	r := chi.NewMux()
	port := ":8080"

	db := map[int64]User{
		1: {
			Username: "John Doe",
			Id:       1,
			Role:     "admin",
			Password: "123456",
		},
	}

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Group(func(r chi.Router) {
		r.Use(jsonMiddleware)
		r.Get("/users/{id:[0-9]+}", handleGetUsers(db))
		r.Post("/users", handlePostUsers(db))
	})

	// r.Get("/users/{id:[0-9]+}", handleGetUsers(db))
	// r.Post("/users", handlePostUsers)

	r.Get("/horario", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Fprintln(w, now)

	})

	fmt.Printf("Server is running on port: %s \n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}

func handleGetUsers(db map[int64]User) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		user, ok := db[id]
		if !ok {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		data, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "Failed to marshal user", http.StatusInternalServerError)
			return

		}
		_, _ = w.Write(data)
	}
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// func handleGetUsers(w http.ResponseWriter, r *http.Request) {

// 	idStr := chi.URLParam(r, "id")
// 	id, err := strconv.ParseInt(idStr, 10, 64)
// 	if err != nil {
// 		http.Error(w, "Invalid ID", http.StatusBadRequest)
// 		return
// 	}

// 	fmt.Fprintln(w, id)
// }

func handlePostUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 1024*1024)
		data, err := io.ReadAll(r.Body)
		if err != nil {

			var maxErr *http.MaxBytesError
			if errors.As(err, &maxErr) {

				sendJson(w, Response{Error: "Request body too large"}, http.StatusRequestEntityTooLarge)

				return
			}

			sendJson(w, Response{Error: "Failed to read body"}, http.StatusInternalServerError)
			return

		}

		var user User
		if err := json.Unmarshal(data, &user); err != nil {
			sendJson(w, Response{Error: "Failed to unmarshal user"}, http.StatusInternalServerError)
			return
		}

		db[user.Id] = user

		w.WriteHeader(http.StatusCreated)

		fmt.Println(string(data))

	}
}
