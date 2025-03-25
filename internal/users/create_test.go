package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	h := handler{db}

	u := User{
		Name:     "gabriel",
		Login:    "gab",
		Password: "abc",
	}

	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(&u)
	if err != nil {
		t.Error(err)
	}

	u.SetPassword(u.Password)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", &b)

	mock.ExpectExec(`insert into "users" ("name", "login", "password", "modified_at")*`).
		WithArgs(u.Name, u.Login, u.Password, u.ModifiedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	h.Create(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("error %v", rr)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}

func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	u, err := New("gabriel", "gab", "1234")
	if err != nil {
		t.Error(err)
	}

	mock.ExpectExec(`insert into "users" ("name", "login", "password", "modified_at")*`).
		WithArgs("gabriel", "gab", u.Password, u.ModifiedAt).WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = Insert(db, u)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
