package repos

import (
	"context"
	"errors"

	"github.com/5822791760/hr/internal/db/postgres/public/model"
	"github.com/5822791760/hr/internal/db/postgres/public/table"
	"github.com/5822791760/hr/pkg/errs"
	"github.com/5822791760/hr/pkg/helpers"
	"github.com/5822791760/hr/pkg/interfaces"

	. "github.com/go-jet/jet/v2/postgres"
)

type Author model.Author

type authorRepo struct {
	clock interfaces.Clock
}

type IAuthorRepo interface {
	NewAuthor(name string, bio string) *Author
	FindAll(ctx context.Context) ([]Author, errs.Err)
	FindOne(ctx context.Context, id int) (*Author, errs.Err)
	Save(ctx context.Context, author *Author) errs.Err
	QueryGetAll(ctx context.Context) ([]QueryAuthorGetAll, errs.Err)
	NameExist(ctx context.Context, name string, id int) (bool, errs.Err)
	Delete(ctx context.Context, id int) errs.Err
}

func NewAuthorRepo(clock interfaces.Clock) authorRepo {
	return authorRepo{clock: clock}
}

func (r authorRepo) NewAuthor(name string, bio string) *Author {
	return &Author{
		Name:      name,
		Bio:       bio,
		CreatedAt: r.clock.Now(),
		UpdatedAt: r.clock.Now(),
	}
}

func (r authorRepo) FindAll(ctx context.Context) ([]Author, errs.Err) {
	db, err := GetDB(ctx)
	if err != nil {
		return []Author{}, err
	}

	q := SELECT(table.Author.AllColumns).FROM(table.Author)

	authors := []Author{}

	if xerr := q.QueryContext(ctx, db, &authors); xerr != nil {
		return []Author{}, errs.NewInternalServerErr(xerr)
	}

	return authors, nil
}

func (r authorRepo) FindOne(ctx context.Context, id int) (*Author, errs.Err) {
	db, err := GetDB(ctx)
	if err != nil {
		return &Author{}, err
	}

	q := SELECT(table.Author.AllColumns).FROM(table.Author).WHERE(table.Author.ID.EQ(Int(int64(id))))

	var author Author

	if xerr := q.QueryContext(ctx, db, &author); xerr != nil {
		return &Author{}, errs.NewAuthorNotFoundErr(xerr)
	}

	return &author, nil
}

func (r authorRepo) Save(ctx context.Context, author *Author) errs.Err {
	var s Statement

	db, err := GetDB(ctx)
	if err != nil {
		return err
	}

	if err := r.Validate(ctx, author); err != nil {
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
			return errs.NewInternalServerErr(xerr)
		}
	}

	return nil
}

func (r authorRepo) Delete(ctx context.Context, id int) errs.Err {
	db, err := GetDB(ctx)
	if err != nil {
		return err
	}

	q := table.Author.DELETE().WHERE(table.Author.ID.EQ(Int(int64(id))))

	res, xerr := q.ExecContext(ctx, db)
	if xerr != nil {
		return errs.NewInternalServerErr(xerr)
	}

	effected, xerr := res.RowsAffected()
	if xerr != nil {
		return errs.NewInternalServerErr(xerr)
	}

	if effected < 1 {
		return errs.NewAuthorNotFoundErr(errors.New("no row affected"))
	}

	return nil
}

func (r authorRepo) NameExist(ctx context.Context, name string, id int) (bool, errs.Err) {
	db, err := GetDB(ctx)
	if err != nil {
		return false, err
	}

	cond := AND(table.Author.Name.EQ(String(name)))

	if id != 0 {
		cond = cond.AND(table.Author.ID.NOT_EQ(Int(int64(id))))
	}

	q := helpers.SelectExist().FROM(table.Author).WHERE(cond)

	exist, err := helpers.IsExist(db, q)
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (r authorRepo) Validate(ctx context.Context, author *Author) errs.Err {
	errCtx := errs.NewErrorContext()

	name := author.Name
	if len(name) < 2 {
		errs.AddAuthorInvalidNameLengthContext(errCtx)
	}

	nameExist, err := r.NameExist(ctx, name, int(author.ID))
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

// ========= QueryGetAll =========

type QueryAuthorGetAll struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (r authorRepo) QueryGetAll(ctx context.Context) ([]QueryAuthorGetAll, errs.Err) {
	db, err := GetDB(ctx)
	if err != nil {
		return []QueryAuthorGetAll{}, err
	}

	data := []QueryAuthorGetAll{}

	q := SELECT(
		table.Author.ID.AS("QueryAuthorGetAll.ID"),
		SELECT(table.Author.Name).AS("QueryAuthorGetAll.Name"),
	).FROM(table.Author)

	if xerr := q.QueryContext(ctx, db, &data); xerr != nil {
		return []QueryAuthorGetAll{}, errs.NewInternalServerErr(xerr)
	}

	return data, nil
}

// ========= QueryGetAll =========
