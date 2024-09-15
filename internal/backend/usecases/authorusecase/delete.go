package authorusecase

import (
	"context"

	"github.com/5822791760/hr/pkg/apperr"
)

// ============================== Response =============================

type DeleteResp struct {
	Success bool `json:"success"`
}

// ============================== Usecase ==============================

func (u authorUsecase) Delete(ctx context.Context, id int) (DeleteResp, apperr.Err) {
	err := u.authorRepo.Delete(ctx, id)
	if err != nil {
		return DeleteResp{}, err
	}

	return DeleteResp{
		Success: true,
	}, nil
}
