package authorusecasetest_test

import (
	"context"
	"testing"

	"github.com/5822791760/hr/internal/repos/authorrepo"
	"github.com/5822791760/hr/internal/usecases/authorusecase"
	"github.com/5822791760/hr/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRead, mockWrite := mocks.GetMockAuthorRepo(ctrl)
	ctx := context.TODO()

	authors := []authorrepo.QueryAuthorGetAll{
		{ID: 1, Name: "Author 1"},
		{ID: 2, Name: "Author 2"},
	}

	mockRead.EXPECT().
		QueryGetAll(ctx).
		Return(authors, nil)

	usecase := authorusecase.NewAuthorUseCase(mockRead, mockWrite)
	res, err := usecase.GetAll(ctx)

	assert.NoError(t, err)
	assert.Equal(t, authors, res)
}
