package authorusecase

import (
	"context"

	"github.com/5822791760/hr/internal/backend/repos"
	"github.com/5822791760/hr/pkg/apperr"
)

type authorUsecase struct {
	authorRepo repos.IAuthorRepo
}

type IAuthorUsecase interface {
	Create(ctx context.Context, body CreateAuthorBody) (CreateAuthorResponse, apperr.Err)
	GetAll(ctx context.Context) ([]repos.QueryGetAllAuthor, apperr.Err)
	GetOne(ctx context.Context, id int) (FindOneAuthorResponse, apperr.Err)
	Update(ctx context.Context, id int, body UpdateAuthorBody) (UpdateAuthorResponse, apperr.Err)
	Delete(ctx context.Context, id int) (DeleteAuthorResponse, apperr.Err)
}

func NewAuthorUseCase(authorRepo repos.IAuthorRepo) authorUsecase {
	return authorUsecase{authorRepo: authorRepo}
}
