package folders

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetRootSubFolders(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "parent_id", "name", "created_at", "modified_at", "deleted"}).
		AddRow(1, nil, "keys", time.Now(), time.Now(), false).
		AddRow(5, nil, "pngs", time.Now(), time.Now(), false)

	mock.ExpectQuery(`select * from "folders" where "parent_id" is null and "deleted" = false`).
		WithArgs().
		WillReturnRows(rows)

	_, err = GetRootSubFolders(db)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
