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
	Create(ctx context.Context, body CreateAuthorBody) (CreateAuthorResponse, errs.Err)
	QueryGetAll(ctx context.Context) ([]repos.QueryAuthorGetAll, errs.Err)
	FindOne(ctx context.Context, id int) (FindOneAuthorResponse, errs.Err)
	Update(ctx context.Context, id int, body UpdateAuthorBody) (UpdateAuthorResponse, errs.Err)
	Delete(ctx context.Context, id int) (DeleteAuthorResponse, errs.Err)
}

func NewAuthorUseCase(authorRepo repos.IAuthorRepo) authorUsecase {
	return authorUsecase{authorRepo: authorRepo}
}

// ============================== QueryGetAll ==============================

func (u authorUsecase) QueryGetAll(ctx context.Context) ([]repos.QueryAuthorGetAll, errs.Err) {
	datas, err := u.authorRepo.QueryGetAll(ctx)
	if err != nil {
		return []repos.QueryAuthorGetAll{}, err
	}

	return datas, nil
}

// ============================== FindOne ==============================

type FindOneAuthorResponse struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Bio  *string `json:"bio"`
}

func (u authorUsecase) FindOne(ctx context.Context, id int) (FindOneAuthorResponse, errs.Err) {
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

// ============================== Create ==============================
type CreateAuthorBody struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

type CreateAuthorResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (u authorUsecase) Create(ctx context.Context, body CreateAuthorBody) (CreateAuthorResponse, errs.Err) {
	author := repos.NewAuthor(body.Name, body.Bio)
	if err := u.authorRepo.Save(ctx, author); err != nil {
		return CreateAuthorResponse{}, err
	}

	return CreateAuthorResponse{
		ID:   int(author.ID),
		Name: author.Name,
	}, nil
}

// ============================== Update ==============================

type UpdateAuthorBody struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

type UpdateAuthorResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (u authorUsecase) Update(ctx context.Context, id int, body UpdateAuthorBody) (UpdateAuthorResponse, errs.Err) {
	author, err := u.authorRepo.FindOne(ctx, id)
	if err != nil {
		return UpdateAuthorResponse{}, err
	}

	author.
		ChangeName(body.Name).
		ChangeBio(body.Bio)

	if err := u.authorRepo.Save(ctx, author); err != nil {
		return UpdateAuthorResponse{}, err
	}

	return UpdateAuthorResponse{
		ID:   int(author.ID),
		Name: author.Name,
	}, nil
}

// ============================== Delete ==============================

type DeleteAuthorResponse struct {
	Success bool `json:"success"`
}

func (u authorUsecase) Delete(ctx context.Context, id int) (DeleteAuthorResponse, errs.Err) {
	err := u.authorRepo.Delete(ctx, id)
	if err != nil {
		return DeleteAuthorResponse{}, err
	}

	return DeleteAuthorResponse{
		Success: true,
	}, nil
}
