package authorusecase

import (
	"context"

	"github.com/5822791760/hr/pkg/apperr"
)

// ============================== Response =============================

type GetOneResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

// ============================== Usecase ==============================

func (u authorUsecase) GetOne(ctx context.Context, id int) (GetOneResp, apperr.Err) {
	author, err := u.authorRepo.FindOne(ctx, id)
	if err != nil {
		return GetOneResp{}, err
	}

	return GetOneResp{
		ID:   int(author.ID),
		Name: author.Name,
		Bio:  author.Bio,
	}, nil
}
