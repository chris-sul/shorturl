package link

import (
	"context"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"chris-sul/shorturl/internal/test"
	"chris-sul/shorturl/database"
)

func TestCreateLinkEntity(t *testing.T) {
	link := database.Link{
		ID: 1,
		UrlCode: "test",
		Destination: "test.com",
	}

	resultLink := createLinkEntity(link)

	assert.Equal(t, resultLink.ID, int32(1))
	assert.Equal(t, resultLink.UrlCode, "test")
	assert.Equal(t, resultLink.Destination, "test.com")
}

func TestRepositoryAll(t *testing.T) {
	db, mock := test.NewDBMock(t)
	defer db.Close()

	repo := NewRepository(db)
	ctx := context.Background()

	// All
	mockRows := sqlmock.NewRows([]string{"id", "url_code", "destination"}).AddRow("1", "test", "test")
	mock.ExpectQuery("select id, url_code, destination from links").WillReturnRows(mockRows)
	links, _ := repo.All(ctx)
	assert.Equal(t, len(links), 1)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
