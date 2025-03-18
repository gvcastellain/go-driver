package files

import (
	"database/sql"
	"net/http"
)

func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {

}

func Insert(db *sql.DB, f *File) (int64, error) {
	stmt := `insert into "files" ("folde_id", "owner_id", "name", "type", "path", "modified_at") VALUES ($1, $2, $3, $4, $5, $6)`

	result, err := db.Exec(stmt, f.FolderID, f.OwnerID, f.Name, f.Type, f.Path, f.ModifiedAt)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}
