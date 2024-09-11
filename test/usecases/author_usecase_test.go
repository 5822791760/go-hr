package usecases_test

import (
	"context"
	"testing"

	"github.com/5822791760/hr/internal/repos"
	"github.com/5822791760/hr/internal/usecases"
	"github.com/5822791760/hr/test/mocks/mock_repos"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAuthorUsecase_QueryGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repos.NewMockIAuthorRepo(ctrl)
	ctx := context.TODO()

	authors := []repos.QueryAuthorGetAll{
		{ID: 1, Name: "Author 1"},
		{ID: 2, Name: "Author 2"},
	}

	mockRepo.EXPECT().
		QueryGetAll(ctx).
		Return(authors, nil)

	usecase := usecases.NewAuthorUseCase(mockRepo)
	res, err := usecase.QueryGetAll(ctx)

	assert.NoError(t, err)
	assert.Equal(t, authors, res)
}

func TestAuthorUsecase_FindOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repos.NewMockIAuthorRepo(ctrl)
	ctx := context.TODO()

	author := &repos.Author{ID: 1, Name: "Author 1"}
	mockRepo.EXPECT().
		FindOne(ctx, 1).
		Return(author, nil)

	usecase := usecases.NewAuthorUseCase(mockRepo)
	res, err := usecase.FindOne(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, usecases.FindOneAuthorResponse{
		ID:   1,
		Name: "Author 1",
		Bio:  "",
	}, res)

}

func TestAuthorUsecase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repos.NewMockIAuthorRepo(ctrl)
	ctx := context.TODO()

	body := usecases.CreateAuthorBody{Name: "New Author", Bio: "New Bio"}
	author := &repos.Author{Name: body.Name, Bio: body.Bio}

	mockRepo.EXPECT().
		NewAuthor(body.Name, body.Bio).
		Return(author)

	mockRepo.EXPECT().
		Save(ctx, author).
		Return(nil)

	usecase := usecases.NewAuthorUseCase(mockRepo)

	res, err := usecase.Create(ctx, body)

	assert.NoError(t, err)
	assert.Equal(t, usecases.CreateAuthorResponse{
		ID:   int(author.ID),
		Name: author.Name,
	}, res)
}

func TestAuthorUsecase_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repos.NewMockIAuthorRepo(ctrl)
	ctx := context.TODO()

	id := 1
	body := usecases.UpdateAuthorBody{Name: "Updated Author", Bio: "Updated Bio"}
	author := &repos.Author{ID: int32(id), Name: "Author 1", Bio: ""}

	mockRepo.EXPECT().
		FindOne(ctx, id).
		Return(author, nil)

	mockRepo.EXPECT().
		Save(ctx, &repos.Author{ID: 1, Name: "Updated Author", Bio: "Updated Bio"}).
		Return(nil)

	usecase := usecases.NewAuthorUseCase(mockRepo)
	res, err := usecase.Update(ctx, id, body)

	assert.NoError(t, err)
	assert.Equal(t, usecases.UpdateAuthorResponse{
		ID:   id,
		Name: "Updated Author",
	}, res)
}

func TestAuthorUsecase_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repos.NewMockIAuthorRepo(ctrl)
	ctx := context.TODO()

	id := 1

	mockRepo.EXPECT().
		Delete(ctx, id).
		Return(nil)

	usecase := usecases.NewAuthorUseCase(mockRepo)
	response, err := usecase.Delete(ctx, id)

	assert.NoError(t, err)
	assert.True(t, response.Success)
}
