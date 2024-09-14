package dbutil

import (
	"github.com/5822791760/hr/pkg/apperr"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

func SelectExist() SelectStatement {
	return SELECT(Int(1))
}

func IsExist(db qrm.DB, statement SelectStatement) (bool, apperr.Err) {
	var data struct {
		Exists bool
	}
	stmt := SELECT(EXISTS(statement).AS("Exists"))
	if xerr := stmt.Query(db, &data); xerr != nil {
		return false, apperr.NewInternalServerErr(xerr)
	}

	return data.Exists, nil
}
