package authorusecase

import (
	"context"

	"github.com/5822791760/hr/pkg/apperr"
)

// ============================== Request ==============================

type CreateAuthorBody struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

// ============================== Response =============================

type CreateAuthorResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ============================== Usecase ==============================

func (u authorUsecase) Create(ctx context.Context, body CreateAuthorBody) (CreateAuthorResponse, apperr.Err) {
	author := u.authorRepo.NewAuthor(body.Name, body.Bio)
	if err := u.authorRepo.Save(ctx, author); err != nil {
		return CreateAuthorResponse{}, err
	}

	return CreateAuthorResponse{
		ID:   int(author.ID),
		Name: author.Name,
	}, nil
}
