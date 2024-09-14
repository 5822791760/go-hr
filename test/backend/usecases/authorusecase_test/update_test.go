package authorusecasetest_test

import (
	"context"
	"testing"

	"github.com/5822791760/hr/internal/backend/repos/authorrepo"
	"github.com/5822791760/hr/internal/backend/usecases/authorusecase"
	"github.com/5822791760/hr/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRead, mockWrite := mocks.GetMockAuthorRepo(ctrl)
	ctx := context.TODO()

	id := 1
	body := authorusecase.UpdateAuthorBody{Name: "Updated Author", Bio: "Updated Bio"}
	author := &authorrepo.Author{ID: int32(id), Name: "Author 1", Bio: ""}

	mockRead.EXPECT().
		FindOne(ctx, id).
		Return(author, nil)

	mockWrite.EXPECT().
		Save(ctx, &authorrepo.Author{ID: 1, Name: "Updated Author", Bio: "Updated Bio"}).
		Return(nil)

	usecase := authorusecase.NewAuthorUseCase(mockRead, mockWrite)
	res, err := usecase.Update(ctx, id, body)

	assert.NoError(t, err)
	assert.Equal(t, authorusecase.UpdateAuthorResponse{
		ID:   id,
		Name: "Updated Author",
	}, res)
}
