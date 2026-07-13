package database

import (
	"context"
	"database/sql"
	"errors"
)

type transaction struct {
	tx *sql.Tx
}

// Commit commits the transaction to the database
func (t *transaction) Commit() error {
	if t.tx == nil {
		return errors.New("transaction is nil")
	}
	return t.tx.Commit()
}

// Rollback rollbacks the transaction from the database
func (t *transaction) Rollback() {
	_ = t.tx.Rollback()
}

// Exec executes a query without returning any rows.
func (t *transaction) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return t.tx.ExecContext(ctx, query, args...)
}

// QueryRow execute a query that is expected to return at most one row.
func (t *transaction) QueryRow(ctx context.Context, query string, args ...any) *sql.Row {
	return t.tx.QueryRowContext(ctx, query, args...)
}

// Query executes a query that returns rows, typically a SELECT.
func (t *transaction) Query(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return t.tx.QueryContext(ctx, query, args...)
}

// Prepare prepares the given query for later execution.
func (t *transaction) Prepare(ctx context.Context, query string) (*sql.Stmt, error) {
	return t.tx.PrepareContext(ctx, query)
}

// Stmt returns a transaction-specific prepared statement from an existing statement.
func (t *transaction) Stmt(ctx context.Context, stmt *sql.Stmt) *sql.Stmt {
	return t.tx.StmtContext(ctx, stmt)
}
