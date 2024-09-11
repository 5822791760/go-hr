package mocks

import (
	"context"
	"time"

	"github.com/5822791760/hr/pkg/helpers"
	"github.com/5822791760/hr/pkg/interfaces"
	"github.com/DATA-DOG/go-sqlmock"
)

func GetDBContext() (context.Context, sqlmock.Sqlmock) {
	ctx := context.Background()
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	ctx = helpers.StoreContextDB(ctx, db)

	return ctx, mock
}

// ===== Fake Clock =====

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

func GetFakeClock() (interfaces.Clock, time.Time) {
	now := time.Now()
	clock := fakeClock{MockTime: now}
	return clock, now
}
