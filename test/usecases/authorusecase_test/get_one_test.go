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

func TestGetOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRead, mockWrite := mocks.GetMockAuthorRepo(ctrl)
	ctx := context.TODO()

	author := &authorrepo.Author{ID: 1, Name: "Author 1"}
	mockRead.EXPECT().
		FindOne(ctx, 1).
		Return(author, nil)

	usecase := authorusecase.NewAuthorUseCase(mockRead, mockWrite)
	res, err := usecase.GetOne(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, authorusecase.FindOneAuthorResponse{
		ID:   1,
		Name: "Author 1",
		Bio:  "",
	}, res)

}
