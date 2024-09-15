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

func TestGetOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockrepo.NewMockIAuthorRepo(ctrl)
	ctx := context.TODO()

	author := &repos.Author{ID: 1, Name: "Author 1"}
	mockRepo.EXPECT().
		FindOne(ctx, 1).
		Return(author, nil)

	usecase := authorusecase.NewAuthorUseCase(mockRepo)
	res, err := usecase.GetOne(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, authorusecase.GetOneResp{
		ID:   1,
		Name: "Author 1",
		Bio:  "",
	}, res)

}
