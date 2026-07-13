package database

import (
	"context"
	"database/sql"
)

// ConnectionPool is the interface for database connection pools
type ConnectionPool interface {
	// NewTransaction returns a new transaction from the database connection
	NewTransaction(ctx context.Context, opts ...TxOption) (Transaction, error)
	// CloseConnections closes all database connections in the pool.
	CloseConnections()
	// GetDatabase returns a database connection
	GetDatabase(value ...string) *sql.DB
}

// Transaction is the interface for database transactions
type Transaction interface {
	// Exec executes a query without returning any rows.
	Exec(ctx context.Context, query string, args ...any) (sql.Result, error)
	// QueryRow execute a query that is expected to return at most one row.
	QueryRow(ctx context.Context, query string, args ...any) *sql.Row
	// Query executes a query that returns rows, typically a SELECT.
	Query(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	// Prepare prepares the given query for later execution.
	Prepare(ctx context.Context, query string) (*sql.Stmt, error)
	// Stmt returns a transaction-specific prepared statement from an existing statement.
	Stmt(ctx context.Context, stmt *sql.Stmt) *sql.Stmt
	// Commit commits the transaction to the database
	Commit() error
	// Rollback rollbacks the transaction from the database.
	Rollback()
}
