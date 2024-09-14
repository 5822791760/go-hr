package authorusecase

import (
	"context"

	"github.com/5822791760/hr/internal/backend/repos/authorrepo"
	"github.com/5822791760/hr/pkg/apperr"
)

type authorUsecase struct {
	authorReadRepo  authorrepo.IReadRepo
	authorWriteRepo authorrepo.IWriteRepo
}

type IAuthorUsecase interface {
	Create(ctx context.Context, body CreateAuthorBody) (CreateAuthorResponse, apperr.Err)
	GetAll(ctx context.Context) ([]authorrepo.QueryGetAll, apperr.Err)
	GetOne(ctx context.Context, id int) (FindOneAuthorResponse, apperr.Err)
	Update(ctx context.Context, id int, body UpdateAuthorBody) (UpdateAuthorResponse, apperr.Err)
	Delete(ctx context.Context, id int) (DeleteAuthorResponse, apperr.Err)
}

func NewAuthorUseCase(authorReadRepo authorrepo.IReadRepo, authorWriteRepo authorrepo.IWriteRepo) authorUsecase {
	return authorUsecase{authorWriteRepo: authorWriteRepo, authorReadRepo: authorReadRepo}
}
