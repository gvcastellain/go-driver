package folders

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/gvcastellain/go-driver/internal/auth"
)

type handler struct {
	db *sql.DB
}

func SetRoutes(r chi.Router, db *sql.DB) {
	h := handler{db}

	r.Group(func(r chi.Router) {
		r.Use(auth.Validate)

		r.Post("/", h.Create)
		r.Put("/{id}", h.Modify)
		r.Delete("/{id}", h.Delete)
		r.Get("/{id}", h.Get)
		r.Get("/", h.List)
	})
}
