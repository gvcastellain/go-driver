package users

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi/v5"
)

func TestModify(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	h := handler{db}

	u := User{
		ID:         1,
		Name:       "gabriel",
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}

	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(&u)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/{id}", &b)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	mock.ExpectExec(`update "users" set "name" = $1, "modified_at" = $2 where id = $3`).
		WithArgs(u.Name, AnyTime{}, u.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	rows := sqlmock.NewRows([]string{"id", "name", "login", "password", "created_at", "modified_at", "deleted", "last_login"}).
		AddRow(1, "gabriel", "gabrieltops29@gmail.com", "123456", time.Now(), time.Now(), false, time.Now())

	mock.ExpectQuery(`select * from "users" where id = $1`).
		WithArgs().
		WillReturnRows(rows)

	h.Modify(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("error %v", rr)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}

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
