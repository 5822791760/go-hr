package postgres

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"time"

	"github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrations embed.FS

func ConnectDB(ctx context.Context, connection string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(30 * time.Minute)

	// Ensure the database connection is available
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}

	postgres.SetQueryLogger(func(ctx context.Context, queryInfo postgres.QueryInfo) {
		// sql, args := queryInfo.Statement.Sql()
		// fmt.Printf("- SQL: %s Args: %v \n\n", sql, args)
		fmt.Printf("\n++++++++++++++++++++++++++++++++\n")
		fmt.Printf("%s \n", queryInfo.Statement.DebugSql())

		// Depending on how the statement is executed, RowsProcessed is:
		//   - Number of rows returned for Query() and QueryContext() methods
		//   - RowsAffected() for Exec() and ExecContext() methods
		//   - Always 0 for Rows() method.
		fmt.Printf("- Rows processed: %d\n", queryInfo.RowsProcessed)
		fmt.Printf("- Duration %s\n", queryInfo.Duration.String())
		fmt.Printf("- Execution error: %v\n", queryInfo.Err)

		callerFile, callerLine, callerFunction := queryInfo.Caller()
		fmt.Printf("- Caller file: %s, line: %d, function: %s\n\n", callerFile, callerLine, callerFunction)
		fmt.Printf("++++++++++++++++++++++++++++++++\n\n")
	})

	return db, nil

}
