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
	Create(ctx context.Context, body CreateBody) (CreateResp, apperr.Err)
	GetAll(ctx context.Context) ([]repos.QueryGetAllAuthor, apperr.Err)
	GetOne(ctx context.Context, id int) (GetOneResp, apperr.Err)
	Update(ctx context.Context, id int, body UpdateBody) (UpdateResp, apperr.Err)
	Delete(ctx context.Context, id int) (DeleteResp, apperr.Err)
}

func NewAuthorUseCase(authorRepo repos.IAuthorRepo) authorUsecase {
	return authorUsecase{authorRepo: authorRepo}
}
