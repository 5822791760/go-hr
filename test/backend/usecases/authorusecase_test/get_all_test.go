package authorusecasetest_test

import (
	"context"
	"testing"

	"github.com/5822791760/hr/internal/backend/repos"
	"github.com/5822791760/hr/internal/backend/usecases/authorusecase"
	"github.com/5822791760/hr/test/mocks/mockrepo"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockrepo.NewMockIAuthorRepo(ctrl)
	ctx := context.TODO()

	authors := []repos.QueryGetAllAuthor{
		{ID: 1, Name: "Author 1"},
		{ID: 2, Name: "Author 2"},
	}

	mockRepo.EXPECT().
		QueryGetAll(ctx).
		Return(authors, nil)

	usecase := authorusecase.NewAuthorUseCase(mockRepo)
	res, err := usecase.GetAll(ctx)

	assert.NoError(t, err)
	assert.Equal(t, authors, res)
}
