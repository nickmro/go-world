package mock

import "github.com/nickmro/world-app/world/db"

type DB struct {
	BeginFn func() (db.Tx, error)
	Tx *Tx
}

// NewDB returns a default mock DB.
func NewDB() *DB {
	database := &DB{
		Tx: NewTx(),
	}

	database.BeginFn = func() (db.Tx, error) {
		return database.Tx, nil
	}

	return database
}

func (db *DB) Begin() (db.Tx, error) {
	return db.BeginFn()
}
