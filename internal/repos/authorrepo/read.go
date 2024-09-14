package authorrepo

import (
	"context"

	"github.com/5822791760/hr/internal/db/schema/hr/public/table"
	"github.com/5822791760/hr/pkg/apperr"
	"github.com/5822791760/hr/pkg/coreutil"
	"github.com/5822791760/hr/pkg/dbutil"
	. "github.com/go-jet/jet/v2/postgres"
)

type readRepo struct{}

func NewReadRepo() readRepo {
	return readRepo{}
}

type IReadRepo interface {
	FindAll(ctx context.Context) ([]Author, apperr.Err)
	FindOne(ctx context.Context, id int) (*Author, apperr.Err)
	QueryGetAll(ctx context.Context) ([]QueryAuthorGetAll, apperr.Err)
	NameExist(ctx context.Context, name string, id int) (bool, apperr.Err)
	Validate(ctx context.Context, author *Author) apperr.Err
}

func (r readRepo) FindAll(ctx context.Context) ([]Author, apperr.Err) {
	db, err := coreutil.GetDB(ctx)
	if err != nil {
		return []Author{}, err
	}

	q := SELECT(table.Author.AllColumns).FROM(table.Author)

	authors := []Author{}

	if xerr := q.QueryContext(ctx, db, &authors); xerr != nil {
		return []Author{}, apperr.NewInternalServerErr(xerr)
	}

	return authors, nil
}

func (r readRepo) FindOne(ctx context.Context, id int) (*Author, apperr.Err) {
	db, err := coreutil.GetDB(ctx)
	if err != nil {
		return &Author{}, err
	}

	q := SELECT(table.Author.AllColumns).FROM(table.Author).WHERE(table.Author.ID.EQ(Int(int64(id))))

	var author Author

	if xerr := q.QueryContext(ctx, db, &author); xerr != nil {
		return &Author{}, apperr.NewAuthorNotFoundErr(xerr)
	}

	return &author, nil
}

func (r readRepo) NameExist(ctx context.Context, name string, id int) (bool, apperr.Err) {
	db, err := coreutil.GetDB(ctx)
	if err != nil {
		return false, err
	}

	cond := AND(table.Author.Name.EQ(String(name)))

	if id != 0 {
		cond = cond.AND(table.Author.ID.NOT_EQ(Int(int64(id))))
	}

	q := dbutil.SelectExist().FROM(table.Author).WHERE(cond)

	exist, err := dbutil.IsExist(db, q)
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (r readRepo) Validate(ctx context.Context, author *Author) apperr.Err {
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

// ========= QueryGetAll =========

type QueryAuthorGetAll struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (r readRepo) QueryGetAll(ctx context.Context) ([]QueryAuthorGetAll, apperr.Err) {
	db, err := coreutil.GetDB(ctx)
	if err != nil {
		return []QueryAuthorGetAll{}, err
	}

	data := []QueryAuthorGetAll{}

	q := SELECT(
		table.Author.ID.AS("QueryAuthorGetAll.ID"),
		SELECT(table.Author.Name).AS("QueryAuthorGetAll.Name"),
	).FROM(table.Author)

	if xerr := q.QueryContext(ctx, db, &data); xerr != nil {
		return []QueryAuthorGetAll{}, apperr.NewInternalServerErr(xerr)
	}

	return data, nil
}

// ========= QueryGetAll =========
