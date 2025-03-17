package files

import "database/sql"

func List(db *sql.DB, folderID int64) ([]File, error) {
	stmt := `select * from "files" where "folder_id" = $1 and "deleted" = false`

	rows, err := db.Query(stmt, folderID)
	if err != nil {
		return nil, err
	}

	files := make([]File, 0)
	for rows.Next() {
		var f File

		err := rows.Scan(&f.ID, &f.FolderID, &f.OwnerID, &f.Name, &f.Type, &f.Path, &f.CreatedAt, &f.ModifiedAt, &f.Deleted)
		if err != nil {
			continue
		}

		files = append(files, f)
	}

	return files, nil
}

func ListRoot(db *sql.DB) ([]File, error) {
	stmt := `select * from "files" where "folder_id" is null and "deleted" = false`

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}

	files := make([]File, 0)
	for rows.Next() {
		var f File

		err := rows.Scan(&f.ID, &f.FolderID, &f.OwnerID, &f.Name, &f.Type, &f.Path, &f.CreatedAt, &f.ModifiedAt, &f.Deleted)
		if err != nil {
			continue
		}

		files = append(files, f)
	}

	return files, nil
}
