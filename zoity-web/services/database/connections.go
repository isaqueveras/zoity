// Package database provides database connection pooling and transaction management.
package database

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"zoity/core"
)

const Default string = "zoity"

type connections struct {
	databases map[string]*sql.DB
}

// OpenConnections creates a new connection pool with the given databases
func OpenConnections(c *core.App) ConnectionPool {
	var (
		pool = make(map[string]*sql.DB)
		err  error
	)

	for idx := range c.Config.Databases {
		if pool[c.Config.Databases[idx].Nick], err = sql.Open("sqlite3", "zoity.db"); err != nil {
			slog.Error("error opening database", "name", c.Config.Databases[idx].Nick, "error", err)
			continue
		}

		pool[c.Config.Databases[idx].Nick].SetMaxOpenConns(c.Config.Databases[idx].MaxConn)
		pool[c.Config.Databases[idx].Nick].SetMaxIdleConns(c.Config.Databases[idx].MaxIdle)
		pool[c.Config.Databases[idx].Nick].SetConnMaxIdleTime(5 * time.Minute)
		pool[c.Config.Databases[idx].Nick].SetConnMaxLifetime(time.Minute * 1)

		if err = pool[c.Config.Databases[idx].Nick].Ping(); err != nil {
			slog.Error("error ping database", "name", c.Config.Databases[idx].Nick, "error", err)
			continue
		}
	}

	return &connections{databases: pool}
}

// CloseConnections closes all connections in the pool
func (c *connections) CloseConnections() {
	for _, db := range c.databases {
		if db == nil {
			continue
		}
		_ = db.Close()
	}
}

// GetDatabase obtem o banco de dados informando um nome ou retorna o padrão
func (c *connections) GetDatabase(value ...string) *sql.DB {
	name := Default
	if len(value) > 0 {
		name = value[0]
	}

	if db, ok := c.databases[name]; ok && db != nil {
		return db
	}

	return nil
}

// NewTransaction starts a new transaction in the database
func (c *connections) NewTransaction(ctx context.Context, opts ...TxOption) (Transaction, error) {
	options := defaultTxOptions()
	for _, opt := range opts {
		opt(&options)
	}

	txOptions := &sql.TxOptions{
		Isolation: options.Isolation,
		ReadOnly:  options.ReadOnly,
	}

	tx, err := c.GetDatabase(options.Database).BeginTx(ctx, txOptions)
	if err != nil {
		return nil, err
	}

	return &transaction{tx: tx}, nil
}
