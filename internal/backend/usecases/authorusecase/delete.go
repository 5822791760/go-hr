package authorusecase

import (
	"context"

	"github.com/5822791760/hr/pkg/apperr"
)

// ============================== Response =============================

type DeleteAuthorResponse struct {
	Success bool `json:"success"`
}

// ============================== Usecase ==============================

func (u authorUsecase) Delete(ctx context.Context, id int) (DeleteAuthorResponse, apperr.Err) {
	err := u.authorWriteRepo.Delete(ctx, id)
	if err != nil {
		return DeleteAuthorResponse{}, err
	}

	return DeleteAuthorResponse{
		Success: true,
	}, nil
}
