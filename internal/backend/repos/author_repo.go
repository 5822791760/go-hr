package repos

import (
	"context"
	"errors"

	"github.com/5822791760/hr/internal/backend/db/schema/hr/public/model"
	"github.com/5822791760/hr/internal/backend/db/schema/hr/public/table"
	"github.com/5822791760/hr/pkg/apperr"
	"github.com/5822791760/hr/pkg/coreutil"
	"github.com/5822791760/hr/pkg/dbutil"
	q "github.com/go-jet/jet/v2/postgres"
)

type Author model.Author

type authorRepo struct {
	clock coreutil.Clock
}

type IAuthorRepo interface {
	NewAuthor(name string, bio string) *Author
	FindAll(ctx context.Context) ([]Author, apperr.Err)
	FindOne(ctx context.Context, id int) (*Author, apperr.Err)
	NameExist(ctx context.Context, name string, id int) (bool, apperr.Err)
	Validate(ctx context.Context, author *Author) apperr.Err
	Save(ctx context.Context, author *Author) apperr.Err
	Delete(ctx context.Context, id int) apperr.Err
	QueryGetAll(ctx context.Context) ([]QueryGetAllAuthor, apperr.Err)
}

func NewAuthorRepo(clock coreutil.Clock) authorRepo {
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

// =================== Read ===================

func (r authorRepo) FindAll(ctx context.Context) ([]Author, apperr.Err) {
	db, err := coreutil.GetDB(ctx)
	if err != nil {
		return []Author{}, err
	}

	stmt := q.SELECT(table.Author.AllColumns).FROM(table.Author)

	authors := []Author{}

	if xerr := stmt.QueryContext(ctx, db, &authors); xerr != nil {
		return []Author{}, apperr.NewInternalServerErr(xerr)
	}

	return authors, nil
}

func (r authorRepo) FindOne(ctx context.Context, id int) (*Author, apperr.Err) {
	db, err := coreutil.GetDB(ctx)
	if err != nil {
		return &Author{}, err
	}

	stmt := q.SELECT(table.Author.AllColumns).FROM(table.Author).WHERE(table.Author.ID.EQ(q.Int(int64(id))))

	var author Author

	if xerr := stmt.QueryContext(ctx, db, &author); xerr != nil {
		return &Author{}, apperr.NewAuthorNotFoundErr(xerr)
	}

	return &author, nil
}

func (r authorRepo) NameExist(ctx context.Context, name string, id int) (bool, apperr.Err) {
	db, err := coreutil.GetDB(ctx)
	if err != nil {
		return false, err
	}

	cond := q.AND(table.Author.Name.EQ(q.String(name)))

	if id != 0 {
		cond = cond.AND(table.Author.ID.NOT_EQ(q.Int(int64(id))))
	}

	stmt := dbutil.SelectExist().FROM(table.Author).WHERE(cond)

	exist, err := dbutil.IsExist(db, stmt)
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (r authorRepo) Validate(ctx context.Context, author *Author) apperr.Err {
	errCtx := apperr.NewErrorContext()

	name := author.Name
	if len(name) < 2 {
		apperr.AddAuthorInvalidNameLengthContext(errCtx)
	}

	nameExist, err := r.NameExist(ctx, name, int(author.ID))
	if err != nil {
		return err
	}

	if nameExist {
		apperr.AddAuthorNameAlreadyExistContext(errCtx)
	}

	if len(errCtx) > 0 {
		return apperr.NewAuthorValidateErr(errCtx)
	}

	return nil
}

// =================== Write ===================

func (r authorRepo) Save(ctx context.Context, author *Author) apperr.Err {
	var stmt q.Statement

	db, err := coreutil.GetDB(ctx)
	if err != nil {
		return err
	}

	if err := r.Validate(ctx, author); err != nil {
		return err
	}

	author.UpdatedAt = r.clock.Now()

	if author.ID == 0 {
		author.CreatedAt = r.clock.Now()

		stmt = table.Author.
			INSERT(table.Author.Name, table.Author.Bio, table.Author.CreatedAt, table.Author.UpdatedAt).
			MODEL(author).
			RETURNING(table.Author.AllColumns)

	} else {
		stmt = table.Author.
			UPDATE(table.Author.Name, table.Author.Bio, table.Author.UpdatedAt).
			MODEL(author).
			WHERE(table.Author.ID.EQ(q.Int(int64(author.ID)))).
			RETURNING(table.Author.AllColumns)
	}

	if stmt != nil {
		if xerr := stmt.QueryContext(ctx, db, author); xerr != nil {
			return apperr.NewInternalServerErr(xerr)
		}
	}

	return nil
}

func (r authorRepo) Delete(ctx context.Context, id int) apperr.Err {
	db, err := coreutil.GetDB(ctx)
	if err != nil {
		return err
	}

	stmt := table.Author.DELETE().WHERE(table.Author.ID.EQ(q.Int(int64(id))))

	res, xerr := stmt.ExecContext(ctx, db)
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

// =================== QueryGetAllAuthor ===================

type QueryGetAllAuthor struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

func (r authorRepo) QueryGetAll(ctx context.Context) ([]QueryGetAllAuthor, apperr.Err) {
	db, err := coreutil.GetDB(ctx)
	if err != nil {
		return []QueryGetAllAuthor{}, err
	}

	data := []QueryGetAllAuthor{}

	stmt := q.SELECT(
		table.Author.ID.AS("QueryGetAllAuthor.ID"),
		table.Author.Name.AS("QueryGetAllAuthor.Name"),
		table.Author.Bio.AS("QueryGetAllAuthor.Bio"),
	).FROM(table.Author)

	if xerr := stmt.QueryContext(ctx, db, &data); xerr != nil {
		return []QueryGetAllAuthor{}, apperr.NewInternalServerErr(xerr)
	}

	return data, nil
}
