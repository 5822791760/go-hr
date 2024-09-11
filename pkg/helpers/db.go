package helpers

import (
	"context"
	"database/sql"

	"github.com/5822791760/hr/pkg/errs"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type Transactionable interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

func SelectExist() SelectStatement {
	return SELECT(Int(1))
}

func IsExist(db qrm.DB, statement SelectStatement) (bool, errs.Err) {
	var data struct {
		Exists bool
	}
	stmt := SELECT(EXISTS(statement).AS("Exists"))
	if xerr := stmt.Query(db, &data); xerr != nil {
		return false, errs.NewInternalServerErr(xerr)
	}

	return data.Exists, nil
}
