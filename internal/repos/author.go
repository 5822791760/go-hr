package repos

import (
	"context"
	"time"

	"github.com/5822791760/hr/internal/db/postgres/public/model"
	"github.com/5822791760/hr/pkg/errs"
)

type Author model.Author

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
	errContexts := errs.NewErrorContext()
	name := author.Name

	if len(name) < 2 {
		errContexts = append(errContexts, errs.NewAuthorInvalidNameLengthContext())
	}

	nameExist, err := repo.NameExist(ctx, name, int(author.ID))
	if err != nil {
		return err
	}

	if nameExist {
		errContexts = append(errContexts, errs.NewAuthorNameAlreadyExistContext())
	}

	if len(errContexts) > 0 {
		return errs.NewAuthorValidateErr(errContexts)
	}

	return nil
}
