package authorusecase

import (
	"context"

	"github.com/5822791760/hr/pkg/apperr"
)

// ============================== Response =============================

type FindOneAuthorResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

// ============================== Usecase ==============================

func (u authorUsecase) GetOne(ctx context.Context, id int) (FindOneAuthorResponse, apperr.Err) {
	author, err := u.authorRepo.FindOne(ctx, id)
	if err != nil {
		return FindOneAuthorResponse{}, err
	}

	return FindOneAuthorResponse{
		ID:   int(author.ID),
		Name: author.Name,
		Bio:  author.Bio,
	}, nil
}
