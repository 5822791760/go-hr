package helpers

import (
	"context"
	"database/sql"

	"github.com/5822791760/hr/pkg/errs"
	"github.com/go-jet/jet/v2/qrm"
)

var coreDB *sql.DB

type contextDBKey string

const DBKey contextDBKey = "DBKey"

func InitCoreDB(db *sql.DB) {
	coreDB = db
}

func StoreContextDB(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, DBKey, coreDB)
	return ctx
}

func StartTransaction(ctx context.Context) (context.Context, errs.Err) {
	tx, xerr := coreDB.BeginTx(ctx, nil)
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

func GetContextDB(ctx context.Context) qrm.DB {
	db, ok := ctx.Value(DBKey).(qrm.DB)
	if !ok {
		return nil
	}

	return db
}
