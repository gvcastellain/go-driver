package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func (h *handler) List(rw http.ResponseWriter, r *http.Request) {
	us, err := SelectAll(h.db)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(us)
}

func SelectAll(db *sql.DB) ([]User, error) {
	stmt := `select * from "users" where deleted = false`

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}

	us := make([]User, 0)

	for rows.Next() {
		var u User

		err := rows.Scan(&u.ID, &u.Name, &u.Login, &u.Password, &u.CreatedAt, &u.ModifiedAt, &u.Deleted, &u.LastLogin)
		if err != nil {
			continue
		}

		us = append(us, u)
	}

	return us, nil
}
