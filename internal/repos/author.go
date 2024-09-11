package repos

import (
	"context"
	"time"

	"github.com/5822791760/hr/internal/db/postgres/public/model"
	"github.com/5822791760/hr/pkg/errs"
)

type Author model.Author

type IAuthor interface {
	ChangeName(name string) *Author
	ChangeBio(bio string) *Author
	LatestUpdate() *Author
	Validate(ctx context.Context, repo authorRepo) errs.Err
}

func NewAuthor(name string, bio string) *Author {
	return &Author{
		Name:      name,
		Bio:       &bio,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (author *Author) ChangeName(name string) *Author {
	if name == author.Name {
		return author
	}

	author.Name = name
	return author
}

func (author *Author) ChangeBio(bio string) *Author {
	if bio == *author.Bio {
		return author
	}

	author.Bio = &bio
	return author
}

func (author *Author) LatestUpdate() *Author {
	author.UpdatedAt = time.Now()
	return author
}

func (author *Author) Validate(ctx context.Context, repo authorRepo) errs.Err {
	errCtx := errs.NewErrorContext()
	name := author.Name

	if len(name) < 2 {
		errs.AddAuthorInvalidNameLengthContext(errCtx)
	}

	nameExist, err := repo.NameExist(ctx, name, int(author.ID))
	if err != nil {
		return err
	}

	if nameExist {
		errs.AddAuthorNameAlreadyExistContext(errCtx)
	}

	if len(errCtx) > 0 {
		return errs.NewAuthorValidateErr(errCtx)
	}

	return nil
}
