package authorusecasetest_test

import (
	"context"
	"testing"

	"github.com/5822791760/hr/internal/backend/usecases/authorusecase"
	"github.com/5822791760/hr/test/mocks/mockrepo"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockrepo.NewMockIAuthorRepo(ctrl)
	ctx := context.TODO()

	id := 1

	mockRepo.EXPECT().
		Delete(ctx, id).
		Return(nil)

	usecase := authorusecase.NewAuthorUseCase(mockRepo)
	response, err := usecase.Delete(ctx, id)

	assert.NoError(t, err)
	assert.True(t, response.Success)
}
