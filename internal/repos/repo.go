package repos

import (
	"context"

	"github.com/5822791760/hr/pkg/helpers"
	"github.com/go-jet/jet/v2/qrm"
)

func GetDB(ctx context.Context, db qrm.DB) qrm.DB {
	ctxDB := helpers.GetContextDB(ctx)
	if ctxDB != nil {
		return ctxDB
	}

	return db
}
