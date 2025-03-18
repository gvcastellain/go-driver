package files

import "database/sql"

func Get(db *sql.DB, id int64) (*File, error) {
	stmt := `select * from "files" where id = $1`

	row := db.QueryRow(stmt, id)

	var f File
	err := row.Scan(&f.ID, &f.FolderID, &f.OwnerID, &f.Name, &f.Type,
		&f.Path, &f.CreatedAt, &f.ModifiedAt, &f.Deleted)

	if err != nil {
		return nil, err
	}

	return &f, nil
}
