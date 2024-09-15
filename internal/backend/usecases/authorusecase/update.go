package authorusecase

import (
	"context"

	"github.com/5822791760/hr/pkg/apperr"
)

// ============================== Request ==============================

type UpdateAuthorBody struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

// ============================== Response =============================

type UpdateAuthorResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ============================== Usecase ==============================

func (u authorUsecase) Update(ctx context.Context, id int, body UpdateAuthorBody) (UpdateAuthorResponse, apperr.Err) {
	author, err := u.authorRepo.FindOne(ctx, id)
	if err != nil {
		return UpdateAuthorResponse{}, err
	}

	author.Name = body.Name
	author.Bio = body.Bio

	if err := u.authorRepo.Save(ctx, author); err != nil {
		return UpdateAuthorResponse{}, err
	}

	return UpdateAuthorResponse{
		ID:   int(author.ID),
		Name: author.Name,
	}, nil
}
