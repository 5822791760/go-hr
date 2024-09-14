package repos_test

import (
	"testing"

	"github.com/5822791760/hr/internal/backend/repos/authorrepo"
	"github.com/5822791760/hr/test/mocks"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAuthorRepo_FindAll(t *testing.T) {
	// Arrange
	ctx, mockDB := mocks.GetDBContext()

	rows := sqlmock.NewRows([]string{"id", "name", "bio"}).AddRow(0, "test", "test")
	mockDB.ExpectQuery(`
		SELECT
			author.id AS "author.id",
			author.name AS "author.name",
			author.bio AS "author.bio",
			author.created_at AS "author.created_at",
			author.updated_at AS "author.updated_at"
		FROM
			public.author;`).
		WillReturnRows(rows)

	// Act
	repo := authorrepo.NewReadRepo()
	_, err := repo.FindAll(ctx)

	// Assert
	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestAuthorRepo_FindOne(t *testing.T) {
	// Arrange
	ctx, mockDB := mocks.GetDBContext()

	rows := sqlmock.NewRows([]string{"id", "name", "bio"}).AddRow(0, "test", "test")
	id := 0

	mockDB.ExpectQuery(`
		SELECT
			author.id AS "author.id",
			author.name AS "author.name",
			author.bio AS "author.bio",
			author.created_at AS "author.created_at",
			author.updated_at AS "author.updated_at"
		FROM
			public.author
		WHERE
			author.id = $1;`).
		WithArgs(id).
		WillReturnRows(rows)

	// Act
	repo := authorrepo.NewReadRepo()
	_, err := repo.FindOne(ctx, id)

	// Assert
	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}
