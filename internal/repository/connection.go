package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

// Connection is an interface that defines common query operations.
type Connection interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

func spanWithQuery(
	ctx context.Context,
	query string,
	args ...interface{},
) context.Context {
	queryWithArgs := interpolateQuery(query, args...)
	log.Printf("Executing query: %s", queryWithArgs)
	return ctx
}

func interpolateQuery(query string, args ...interface{}) string {
	queryWithArgs := query
	for _, arg := range args {
		argStr := fmt.Sprintf("%v", arg)
		queryWithArgs = strings.Replace(queryWithArgs, "$", argStr, 1)
	}
	return queryWithArgs
}
