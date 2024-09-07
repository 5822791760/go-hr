package usecases_test

import (
	"context"
	"testing"

	"github.com/5822791760/hr/internal/repos"
	"github.com/5822791760/hr/internal/usecases"
	"github.com/5822791760/hr/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAuthorUsecase_QueryGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIAuthorRepo(ctrl)
	ctx := context.TODO()

	expectedAuthors := []repos.QueryAuthorGetAll{
		{ID: 1, Name: "Author 1"},
		{ID: 2, Name: "Author 2"},
	}

	mockRepo.EXPECT().
		QueryGetAll(ctx).
		Return(expectedAuthors, nil)

	usecase := usecases.NewAuthorUseCase(mockRepo)
	authors, err := usecase.QueryGetAll(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedAuthors, authors)
}

func TestAuthorUsecase_FindOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIAuthorRepo(ctrl)
	ctx := context.TODO()

	expectedAuthor := &repos.Author{ID: 1, Name: "Author 1"}
	mockRepo.EXPECT().
		FindOne(ctx, 1).
		Return(expectedAuthor, nil)

	usecase := usecases.NewAuthorUseCase(mockRepo)
	author, err := usecase.FindOne(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, *expectedAuthor, author)
}

func TestAuthorUsecase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIAuthorRepo(ctrl)
	ctx := context.TODO()

	body := usecases.CreateAuthorBody{Name: "New Author", Bio: "New Bio"}
	expectedAuthor := &repos.Author{Name: body.Name}

	mockRepo.EXPECT().
		Save(ctx, gomock.Any()).
		Return(nil)

	usecase := usecases.NewAuthorUseCase(mockRepo)
	author, err := usecase.Create(ctx, body)

	assert.NoError(t, err)
	assert.Equal(t, *expectedAuthor, author)
}

// func TestAuthorUsecase_Update(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRepo := mocks.NewMockIAuthorRepo(ctrl)
// 	ctx := context.TODO()

// 	id := 1
// 	body := usecases.UpdateAuthorBody{Name: "Updated Author", Bio: "Updated Bio"}
// 	existingAuthor := &repos.Author{ID: int32(id), Name: "Author 1"}

// 	mockRepo.EXPECT().
// 		FindOne(ctx, id).
// 		Return(existingAuthor, nil)

// 	mockRepo.EXPECT().
// 		Save(ctx, existingAuthor).
// 		Return(nil)

// 	usecase := usecases.NewAuthorUseCase(mockRepo)
// 	author, err := usecase.Update(ctx, id, body)

// 	assert.NoError(t, err)
// 	assert.Equal(t, *existingAuthor, author)
// }

// func TestAuthorUsecase_Delete(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRepo := mocks.NewMockIAuthorRepo(ctrl)
// 	ctx := context.TODO()

// 	id := 1

// 	mockRepo.EXPECT().
// 		Delete(ctx, id).
// 		Return(nil)

// 	usecase := usecases.NewAuthorUseCase(mockRepo)
// 	response, err := usecase.Delete(ctx, id)

// 	assert.NoError(t, err)
// 	assert.True(t, response.Success)
// }
