package usecases

import (
	"context"

	"github.com/5822791760/hr/internal/repos"
	"github.com/5822791760/hr/pkg/errs"
)

type authorUsecase struct {
	authorRepo repos.IAuthorRepo
}

type IAuthorUsecase interface {
	QueryGetAll(ctx context.Context) ([]repos.QueryAuthorGetAll, errs.Err)
	FindOne(ctx context.Context, id int) (repos.Author, errs.Err)
}

func NewAuthorUseCase(authorRepo repos.IAuthorRepo) authorUsecase {
	return authorUsecase{authorRepo: authorRepo}
}

func (u authorUsecase) QueryGetAll(ctx context.Context) ([]repos.QueryAuthorGetAll, errs.Err) {
	datas, err := u.authorRepo.QueryGetAll(ctx)
	if err != nil {
		return []repos.QueryAuthorGetAll{}, err
	}

	return datas, nil
}

func (u authorUsecase) FindOne(ctx context.Context, id int) (repos.Author, errs.Err) {
	author, err := u.authorRepo.FindOne(ctx, int64(id))
	if err != nil {
		return repos.Author{}, err
	}

	return author, nil
}
