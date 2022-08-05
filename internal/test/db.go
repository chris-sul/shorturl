package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
	"database/sql"
)

func NewDBMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error making mock db!", err)
	}
	return db, mock
}
