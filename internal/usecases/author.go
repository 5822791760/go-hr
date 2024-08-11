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
	if err := u.Validate(ctx, ValidateAuthorBody{Name: &body.Name}, nil); err != nil {
		return repos.Author{}, err
	}

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
	if err := u.Validate(ctx, ValidateAuthorBody{Name: &body.Name}, &id); err != nil {
		return repos.Author{}, err
	}

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

// ====== Validate =======
type ValidateAuthorBody struct {
	ID   *int
	Name *string
}

func (u authorUsecase) Validate(ctx context.Context, req ValidateAuthorBody, id *int) errs.Err {
	errContexts := errs.NewErrorContext()
	name := req.Name

	if name != nil {
		if len(*name) < 2 {
			errContexts = append(errContexts, errs.NewAuthorInvalidNameLengthContext())
		}

		nameExist, err := u.authorRepo.NameExist(*name, id)
		if err != nil {
			return err
		}

		if nameExist {
			errContexts = append(errContexts, errs.NewAuthorNameAlreadyExistContext())
		}
	}

	if len(errContexts) > 0 {
		return errs.NewAuthorValidateErr(errContexts)
	}

	return nil
}

// ====== Validate =======
