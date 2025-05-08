package users

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/gvcastellain/go-driver/internal/auth"
)

var gloabalHandler handler

type handler struct {
	db *sql.DB
}

func SetRoutes(r chi.Router, db *sql.DB) {
	gloabalHandler = handler{db}

	r.Route("/users", func(r chi.Router) {
		r.Post("/", gloabalHandler.Create) //todo - auth not required

		r.Group(func(r chi.Router) {
			r.Use(auth.Validate)

			r.Put("/{id}", gloabalHandler.Modify)
			r.Delete("/{id}", gloabalHandler.Delete)
			r.Get("/{id}", gloabalHandler.GetByID)
			r.Get("/", gloabalHandler.List)
		})
	})
}
