package folders

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {
	f := new(Folder)

	err := json.NewDecoder(r.Body).Decode(f)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = f.Validade()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := Insert(h.db, f)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	f.ID = id
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(f)
}

func Insert(db *sql.DB, f *Folder) (int64, error) {
	stmt := `insert into "folders" ("parent_id", "name", "modified") values ($1, $2, $3)`

	result, err := db.Exec(stmt, f.ParentID, f.Name, f.ModifiedAt)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}
