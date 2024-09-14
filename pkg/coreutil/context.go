package coreutil

import (
	"context"
	"database/sql"

	"github.com/5822791760/hr/pkg/apperr"
	"github.com/go-jet/jet/v2/qrm"
)

type contextDBKey string

const DBKey contextDBKey = "DBKey"

type Transactionable interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

func GetDB(ctx context.Context) (qrm.DB, apperr.Err) {
	ctxDB, err := GetContextDB(ctx)
	if err != nil {
		return nil, err
	}

	return ctxDB, nil
}

func StoreContextDB(ctx context.Context, db interface{}) context.Context {
	ctx = context.WithValue(ctx, DBKey, db)
	return ctx
}

func StartTransaction(ctx context.Context, db Transactionable) (context.Context, apperr.Err) {
	tx, xerr := db.BeginTx(ctx, nil)
	if xerr != nil {
		return nil, apperr.NewInternalServerErr(xerr)
	}

	ctx = context.WithValue(ctx, DBKey, tx)

	return ctx, nil
}

func GetContextTx(ctx context.Context) (*sql.Tx, apperr.Err) {
	db, ok := ctx.Value(DBKey).(*sql.Tx)
	if !ok {
		return nil, apperr.NewInternalServerErrByString("DB not found in context")
	}

	return db, nil
}

func GetContextDB(ctx context.Context) (qrm.DB, apperr.Err) {
	db, ok := ctx.Value(DBKey).(qrm.DB)
	if !ok {
		return nil, apperr.NewInternalServerErrByString("No DB context stored")
	}

	return db, nil
}
