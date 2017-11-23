package postgresql_test

import (
	"github.com/nickmro/world-app/world/db"
)

// Tx is a test wrapper for postgres.Tx.
type Tx struct {
	db.Tx
}
