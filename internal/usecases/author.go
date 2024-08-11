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
	Create(ctx context.Context, body CreateAuthorBody) (repos.Author, errs.Err)
	QueryGetAll(ctx context.Context) ([]repos.QueryAuthorGetAll, errs.Err)
	FindOne(ctx context.Context, id int) (repos.Author, errs.Err)
	Update(ctx context.Context, id int, body UpdateAuthorBody) (repos.Author, errs.Err)
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
	author, err := u.authorRepo.FindOne(ctx, id)
	if err != nil {
		return repos.Author{}, err
	}

	return *author, nil
}

// ===== Create =======
type CreateAuthorBody struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

func (u authorUsecase) Create(ctx context.Context, body CreateAuthorBody) (repos.Author, errs.Err) {
	author := repos.NewAuthor(body.Name, body.Bio)
	if err := u.authorRepo.Save(ctx, author); err != nil {
		return repos.Author{}, err
	}

	return *author, nil
}

// ===== Create =======

// ===== Update =======
type UpdateAuthorBody struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

func (u authorUsecase) Update(ctx context.Context, id int, body UpdateAuthorBody) (repos.Author, errs.Err) {
	author, err := u.authorRepo.FindOne(ctx, id)
	if err != nil {
		return repos.Author{}, err
	}

	author.
		ChangeName(body.Name).
		ChangeBio(body.Bio)

	if err := u.authorRepo.Save(ctx, author); err != nil {
		return repos.Author{}, err
	}

	return *author, nil
}

// ===== Update =======
