package users

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	mock.ExpectExec(`update "users" set "name" = $1, "modified_at" = $2 where id = $3`).
		WithArgs("Gabriel", AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = Update(db, 1, &User{Name: "Gabriel"})
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
