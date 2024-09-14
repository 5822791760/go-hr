package testutil

import (
	"context"

	"github.com/5822791760/hr/pkg/coreutil"
	"github.com/DATA-DOG/go-sqlmock"
)

func GetDBContext() (context.Context, sqlmock.Sqlmock) {
	ctx := context.Background()
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	ctx = coreutil.StoreContextDB(ctx, db)

	return ctx, mock
}
