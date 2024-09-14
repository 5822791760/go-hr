package dbutil

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
)

type ConnectOptions struct {
	Connection  string
	Pool        int
	MaxLifeTime time.Duration
	Logging     bool
}

func ConnectDB(ctx context.Context, opts ConnectOptions) (*sql.DB, error) {
	if opts.Connection == "" {
		return nil, errors.New("empty connection string")
	}

	if opts.Pool == 0 {
		opts.Pool = 20
	}

	if opts.MaxLifeTime == 0 {
		opts.MaxLifeTime = 30 * time.Minute
	}

	db, err := sql.Open("postgres", opts.Connection)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(opts.Pool)
	db.SetMaxIdleConns(opts.Pool)
	db.SetConnMaxLifetime(opts.MaxLifeTime)

	// Ensure the database Connection is available
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	if opts.Logging {
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
	}

	return db, nil

}
