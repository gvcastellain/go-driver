package users

import (
	"database/sql"
	"net/http"
)

func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {

}

func Insert(db *sql.DB, u *User) (int64, error) {
	stmt := `insert into "users" ("name", "login", "password", "modified_at") VALUES ($1, $2, $3, $4)`

	result, err := db.Exec(stmt, u.Name, u.Login, u.Password, u.ModifiedAt)

	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}
