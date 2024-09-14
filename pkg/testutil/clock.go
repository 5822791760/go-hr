package testutil

import (
	"time"

	"github.com/5822791760/hr/pkg/coreutil"
	"github.com/5822791760/hr/test/mocks/repos/mock_authorrepo"
	"go.uber.org/mock/gomock"
)

type fakeClock struct {
	MockTime time.Time
}

func (c fakeClock) Now() time.Time {
	return c.MockTime
}

func (c fakeClock) After(d time.Duration) <-chan time.Time {
	ch := make(chan time.Time, 1)
	ch <- c.MockTime.Add(d)
	return ch
}

func GetFakeClock() (coreutil.Clock, time.Time) {
	now := time.Now()
	clock := fakeClock{MockTime: now}
	return clock, now
}

func GetMockAuthorRepo(ctrl *gomock.Controller) (*mock_authorrepo.MockIReadRepo, *mock_authorrepo.MockIWriteRepo) {
	mockRead := mock_authorrepo.NewMockIReadRepo(ctrl)
	mockWrite := mock_authorrepo.NewMockIWriteRepo(ctrl)

	return mockRead, mockWrite
}
