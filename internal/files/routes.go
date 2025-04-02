package files

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/gvcastellain/go-driver/internal/auth"
	"github.com/gvcastellain/go-driver/internal/bucket"
	"github.com/gvcastellain/go-driver/internal/queue"
)

type handler struct {
	db     *sql.DB
	bucket *bucket.Bucket
	queue  *queue.Queue
}

func SetRoutes(r chi.Router, db *sql.DB, b *bucket.Bucket, q *queue.Queue) {
	h := handler{db, b, q}

	r.Group(func(r chi.Router) {
		r.Use(auth.Validate) //adds auth-Bearer to request

		r.Post("/", h.Create)
		r.Put("/{id}", h.Modify)
		r.Delete("/{id}", h.Delete)
	})
}
