package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {
	u := new(User) //create as pointer

	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	u.SetPassword(u.Password)

	err = u.Validade()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := Insert(h.db, u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	u.ID = id

	rw.WriteHeader(http.StatusCreated)
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(u)
}

func Insert(db *sql.DB, u *User) (int64, error) {
	stmt := `insert into "users" ("name", "login", "password", "modified_at") VALUES ($1, $2, $3, $4)`

	result, err := db.Exec(stmt, u.Name, u.Login, u.Password, u.ModifiedAt)

	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}
