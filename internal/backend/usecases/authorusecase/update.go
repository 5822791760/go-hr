package authorusecase

import (
	"context"

	"github.com/5822791760/hr/pkg/apperr"
)

// ============================== Request ==============================

type UpdateBody struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

// ============================== Response =============================

type UpdateResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ============================== Usecase ==============================

func (u authorUsecase) Update(ctx context.Context, id int, body UpdateBody) (UpdateResp, apperr.Err) {
	author, err := u.authorRepo.FindOne(ctx, id)
	if err != nil {
		return UpdateResp{}, err
	}

	author.Name = body.Name
	author.Bio = body.Bio

	if err := u.authorRepo.Save(ctx, author); err != nil {
		return UpdateResp{}, err
	}

	return UpdateResp{
		ID:   int(author.ID),
		Name: author.Name,
	}, nil
}
