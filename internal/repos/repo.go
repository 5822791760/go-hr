package repos

import (
	"context"

	"github.com/5822791760/hr/pkg/errs"
	"github.com/5822791760/hr/pkg/helpers"
	"github.com/go-jet/jet/v2/qrm"
)

func GetDB(ctx context.Context) (qrm.DB, errs.Err) {
	ctxDB, err := helpers.GetContextDB(ctx)
	if err != nil {
		return nil, err
	}

	return ctxDB, nil
}
