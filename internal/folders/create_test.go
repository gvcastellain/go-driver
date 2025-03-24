package folders

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	f, err := New("photos", 0)
	if err != nil {
		t.Error(err)
	}

	mock.ExpectExec(`insert into "folders" ("parent_id", "name", "modified") values ($1, $2, $3)`).
		WithArgs(0, f.Name, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = Insert(db, f)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
