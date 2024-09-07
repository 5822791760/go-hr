package repos

import (
	"context"
	"errors"

	"github.com/5822791760/hr/internal/db/postgres/public/table"
	"github.com/5822791760/hr/pkg/errs"
	"github.com/5822791760/hr/pkg/helpers"

	. "github.com/go-jet/jet/v2/postgres"
)

type authorRepo struct{}

type IAuthorRepo interface {
	FindAll(ctx context.Context) ([]*Author, errs.Err)
	FindOne(ctx context.Context, id int) (*Author, errs.Err)
	Save(ctx context.Context, author *Author) errs.Err
	QueryGetAll(ctx context.Context) ([]QueryAuthorGetAll, errs.Err)
	NameExist(ctx context.Context, name string, id int) (bool, errs.Err)
	Delete(ctx context.Context, id int) errs.Err
}

func NewAuthorRepo() authorRepo {
	return authorRepo{}
}

func (r authorRepo) FindAll(ctx context.Context) ([]*Author, errs.Err) {
	db, err := GetDB(ctx)
	if err != nil {
		return []*Author{}, err
	}

	q := SELECT(table.Author.AllColumns).FROM(table.Author)

	authors := []*Author{}

	if xerr := q.QueryContext(ctx, db, authors); xerr != nil {
		return []*Author{}, errs.NewInternalServerErr(xerr)
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
	var insertStmt InsertStatement
	var updateStmt UpdateStatement

	db, err := GetDB(ctx)
	if err != nil {
		return err
	}

	if err := author.Validate(ctx, r); err != nil {
		return err
	}

	author.LatestUpdate()

	if author.ID == 0 {
		insertStmt = table.Author.
			INSERT(table.Author.Name, table.Author.Bio).
			VALUES(author.Name, *author.Bio).
			RETURNING(table.Author.AllColumns)

	} else {
		updateStmt = table.Author.UPDATE(table.Author.AllColumns).MODEL(author).WHERE(table.Author.ID.EQ(Int(int64(author.ID)))).RETURNING(table.Author.AllColumns)
	}

	if insertStmt != nil {
		if xerr := insertStmt.QueryContext(ctx, db, author); xerr != nil {
			return errs.NewInternalServerErr(xerr)
		}
		// query, args := insertStmt.Sql()
		// rows, xerr := db.QueryContext(ctx, query, args...)
		// if xerr != nil {
		// 	return errs.NewInternalServerErr(xerr)
		// }

		// for rows.Next() {
		// 	if xerr := rows.Scan(&author.ID); xerr != nil {
		// 		return errs.NewInternalServerErr(xerr)
		// 	}
		// }
	}

	if updateStmt != nil {
		if xerr := updateStmt.QueryContext(ctx, db, author); xerr != nil {
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
