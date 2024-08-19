package api

import (
	"net/http"
	"user-api/src/db"
	"user-api/src/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(db db.DB) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/api/users", routes.CreateUser(db))
	r.Get("/api/users", routes.FetchUsers(db))
	r.Get("/api/users/{id}", routes.GetUserById(db))
	r.Delete("/api/users/{id}", routes.DeleteUserById(db))
	r.Put("/api/users/{id}", routes.UpdateUserById(db))

	return r
}
