package helpers

import (
	"database/sql"

	"github.com/5822791760/hr/pkg/errs"
	. "github.com/go-jet/jet/v2/postgres"
)

func SelectExist() SelectStatement {
	return SELECT(Int(1))
}

func IsExist(db *sql.DB, statement SelectStatement) (bool, errs.Err) {
	var data struct {
		Exists bool
	}
	stmt := SELECT(EXISTS(statement).AS("Exists"))
	if err := stmt.Query(db, &data); err != nil {
		return false, errs.NewInternalServerErr(err)
	}

	return data.Exists, nil
}
