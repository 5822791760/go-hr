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

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockrepo.NewMockIAuthorRepo(ctrl)
	ctx := context.TODO()

	id := 1
	body := authorusecase.UpdateBody{Name: "Updated Author", Bio: "Updated Bio"}
	author := &repos.Author{ID: int32(id), Name: "Author 1", Bio: ""}

	mockRepo.EXPECT().
		FindOne(ctx, id).
		Return(author, nil)

	mockRepo.EXPECT().
		Save(ctx, &repos.Author{ID: 1, Name: "Updated Author", Bio: "Updated Bio"}).
		Return(nil)

	usecase := authorusecase.NewAuthorUseCase(mockRepo)
	res, err := usecase.Update(ctx, id, body)

	assert.NoError(t, err)
	assert.Equal(t, authorusecase.UpdateResp{
		ID:   id,
		Name: "Updated Author",
	}, res)
}
