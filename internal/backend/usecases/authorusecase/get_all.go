package authorusecase

import (
	"context"

	"github.com/5822791760/hr/internal/backend/repos"
	"github.com/5822791760/hr/pkg/apperr"
)

// ============================== Usecase ==============================

func (u authorUsecase) GetAll(ctx context.Context) ([]repos.QueryGetAllAuthor, apperr.Err) {
	datas, err := u.authorRepo.QueryGetAll(ctx)
	if err != nil {
		return []repos.QueryGetAllAuthor{}, err
	}

	return datas, nil
}
