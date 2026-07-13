package database

import "database/sql"

// TxOptions modela os dados de uma opção de transação
type TxOptions struct {
	Database  string
	ReadOnly  bool
	Isolation sql.IsolationLevel
}

func defaultTxOptions() TxOptions {
	return TxOptions{
		Database:  Default,
		ReadOnly:  false,
		Isolation: sql.LevelDefault,
	}
}

// TxOption define o tipo de opções
type TxOption func(*TxOptions)

// WithIsolation define o nivel de isolação da transação
func WithIsolation(level sql.IsolationLevel) TxOption {
	return func(o *TxOptions) {
		o.Isolation = level
	}
}

// ReadOnly define que a transação será apenas de leitura
func ReadOnly() TxOption {
	return func(o *TxOptions) {
		o.ReadOnly = true
	}
}

// WithDatabase define a opção do nome do banco de dados
func WithDatabase(name string) TxOption {
	return func(o *TxOptions) {
		o.Database = name
	}
}
