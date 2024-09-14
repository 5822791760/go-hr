package authorusecase

import (
	"context"

	"github.com/5822791760/hr/internal/backend/repos/authorrepo"
	"github.com/5822791760/hr/pkg/apperr"
)

// ============================== Usecase ==============================

func (u authorUsecase) GetAll(ctx context.Context) ([]authorrepo.QueryGetAll, apperr.Err) {
	datas, err := u.authorReadRepo.QueryGetAll(ctx)
	if err != nil {
		return []authorrepo.QueryGetAll{}, err
	}

	return datas, nil
}
