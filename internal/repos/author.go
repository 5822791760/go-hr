package repos

import (
	"time"

	"github.com/5822791760/hr/internal/db/postgres/public/model"
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
