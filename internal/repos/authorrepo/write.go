package authorrepo

import (
	"context"
	"errors"

	"github.com/5822791760/hr/internal/db/schema/hr/public/table"
	"github.com/5822791760/hr/pkg/apperr"
	"github.com/5822791760/hr/pkg/coreutil"

	. "github.com/go-jet/jet/v2/postgres"
)

type writeRepo struct {
	readRepo IReadRepo
	clock    coreutil.Clock
}

type IWriteRepo interface {
	NewAuthor(name string, bio string) *Author
	Save(ctx context.Context, author *Author) apperr.Err
	Delete(ctx context.Context, id int) apperr.Err
}

func NewWriteRepo(readRepo IReadRepo, clock coreutil.Clock) writeRepo {
	return writeRepo{readRepo: readRepo, clock: clock}
}

func (r writeRepo) NewAuthor(name string, bio string) *Author {
	return &Author{
		Name:      name,
		Bio:       bio,
		CreatedAt: r.clock.Now(),
		UpdatedAt: r.clock.Now(),
	}
}

func (r writeRepo) Save(ctx context.Context, author *Author) apperr.Err {
	var s Statement

	db, err := coreutil.GetDB(ctx)
	if err != nil {
		return err
	}

	if err := r.readRepo.Validate(ctx, author); err != nil {
		return err
	}

	author.UpdatedAt = r.clock.Now()

	if author.ID == 0 {
		author.CreatedAt = r.clock.Now()

		s = table.Author.
			INSERT(table.Author.Name, table.Author.Bio, table.Author.CreatedAt, table.Author.UpdatedAt).
			MODEL(author).
			RETURNING(table.Author.AllColumns)

	} else {
		s = table.Author.
			UPDATE(table.Author.Name, table.Author.Bio, table.Author.UpdatedAt).
			MODEL(author).
			WHERE(table.Author.ID.EQ(Int(int64(author.ID)))).
			RETURNING(table.Author.AllColumns)
	}

	if s != nil {
		if xerr := s.QueryContext(ctx, db, author); xerr != nil {
			return apperr.NewInternalServerErr(xerr)
		}
	}

	return nil
}

func (r writeRepo) Delete(ctx context.Context, id int) apperr.Err {
	db, err := coreutil.GetDB(ctx)
	if err != nil {
		return err
	}

	q := table.Author.DELETE().WHERE(table.Author.ID.EQ(Int(int64(id))))

	res, xerr := q.ExecContext(ctx, db)
	if xerr != nil {
		return apperr.NewInternalServerErr(xerr)
	}

	effected, xerr := res.RowsAffected()
	if xerr != nil {
		return apperr.NewInternalServerErr(xerr)
	}

	if effected < 1 {
		return apperr.NewAuthorNotFoundErr(errors.New("no row affected"))
	}

	return nil
}
