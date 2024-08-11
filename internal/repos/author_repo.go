package repos

import (
	"context"
	"database/sql"

	"github.com/5822791760/hr/internal/db/postgres/public/table"
	"github.com/5822791760/hr/pkg/errs"

	. "github.com/go-jet/jet/v2/postgres"
)

type authorRepo struct {
	db *sql.DB
}

type IAuthorRepo interface {
	FindAll(ctx context.Context) ([]Author, errs.Err)
	FindOne(ctx context.Context, id int64) (Author, errs.Err)
	Save(ctx context.Context, author *Author) errs.Err
	QueryGetAll(ctx context.Context) ([]QueryAuthorGetAll, errs.Err)
}

func NewAuthorRepo(db *sql.DB) authorRepo {
	return authorRepo{db: db}
}

func (r authorRepo) FindAll(ctx context.Context) ([]Author, errs.Err) {
	q := SELECT(table.Author.AllColumns).FROM(table.Author)

	authors := []Author{}

	if err := q.QueryContext(ctx, r.db, &authors); err != nil {
		return []Author{}, errs.NewInternalServerErr(err)
	}

	return authors, nil
}

func (r authorRepo) FindOne(ctx context.Context, id int64) (Author, errs.Err) {
	q := SELECT(table.Author.AllColumns).FROM(table.Author).WHERE(table.Author.ID.EQ(Int(id)))

	var author Author

	if err := q.QueryContext(ctx, r.db, &author); err != nil {
		return Author{}, errs.NewAuthorNotFoundErr(err)
	}

	return author, nil
}

func (r authorRepo) Save(ctx context.Context, author *Author) errs.Err {
	var insertStmt InsertStatement
	var updateStmt UpdateStatement
	var err errs.Err

	if author.ID == 0 {
		insertStmt = table.Author.INSERT(table.Author.AllColumns).MODEL(author).RETURNING(table.Author.AllColumns)
	} else {
		updateStmt = table.Author.UPDATE(table.Author.AllColumns).MODEL(author).RETURNING(table.Author.AllColumns)
	}

	author.LatestUpdate()

	if insertStmt != nil {
		insertErr := insertStmt.QueryContext(ctx, r.db, author)
		err = errs.NewInternalServerErr(insertErr)
	}

	if updateStmt != nil {
		updateErr := updateStmt.QueryContext(ctx, r.db, author)
		err = errs.NewInternalServerErr(updateErr)
	}

	return err
}

// ========= QueryGetAll =========

type QueryAuthorGetAll struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (r authorRepo) QueryGetAll(ctx context.Context) ([]QueryAuthorGetAll, errs.Err) {
	data := []QueryAuthorGetAll{}

	q := SELECT(
		table.Author.ID.AS("QueryAuthorGetAll.ID"),
		SELECT(table.Author.Name).AS("QueryAuthorGetAll.Name"),
	).FROM(table.Author)

	if err := q.QueryContext(ctx, r.db, &data); err != nil {
		return []QueryAuthorGetAll{}, errs.NewInternalServerErr(err)
	}

	return data, nil
}

// ========= QueryGetAll =========
