package helpers

import (
	"context"
	"database/sql"

	"github.com/5822791760/hr/pkg/errs"
	"github.com/go-jet/jet/v2/qrm"
)

type contextDBKey string

const DBKey contextDBKey = "DBKey"

func StoreContextDB(ctx context.Context, db interface{}) context.Context {
	ctx = context.WithValue(ctx, DBKey, db)
	return ctx
}

func StartTransaction(ctx context.Context, db Transactionable) (context.Context, errs.Err) {
	tx, xerr := db.BeginTx(ctx, nil)
	if xerr != nil {
		return nil, errs.NewInternalServerErr(xerr)
	}

	ctx = context.WithValue(ctx, DBKey, tx)

	return ctx, nil
}

func GetContextTx(ctx context.Context) (*sql.Tx, errs.Err) {
	db, ok := ctx.Value(DBKey).(*sql.Tx)
	if !ok {
		return nil, errs.NewInternalServerErrByString("DB not found in context")
	}

	return db, nil
}

func GetContextDB(ctx context.Context) (qrm.DB, errs.Err) {
	db, ok := ctx.Value(DBKey).(qrm.DB)
	if !ok {
		return nil, errs.NewInternalServerErrByString("No DB context stored")
	}

	return db, nil
}
