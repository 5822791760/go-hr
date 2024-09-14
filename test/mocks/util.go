package mocks

import (
	"github.com/5822791760/hr/test/mocks/repos/mock_authorrepo"
	"go.uber.org/mock/gomock"
)

func GetMockAuthorRepo(ctrl *gomock.Controller) (*mock_authorrepo.MockIReadRepo, *mock_authorrepo.MockIWriteRepo) {
	mockRead := mock_authorrepo.NewMockIReadRepo(ctrl)
	mockWrite := mock_authorrepo.NewMockIWriteRepo(ctrl)

	return mockRead, mockWrite
}
