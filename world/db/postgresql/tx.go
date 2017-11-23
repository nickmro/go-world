package postgresql

import (
	"github.com/drinkin/di/lg"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
)

// Tx wraps a runner.Tx instance.
type Tx struct {
	*runner.Tx
	logger *lg.Logger
}

// Commit commits a database transaction.
func (tx Tx) Commit() error {
	return tx.Tx.Commit()
}

// Rollback rolls back a transaction.
func (tx Tx) Rollback() error {
	return tx.Tx.Rollback()
}

// AutoRollback automatically rolls back a transaction.
func (tx Tx) AutoRollback() error {
	return tx.Tx.AutoRollback()
}
