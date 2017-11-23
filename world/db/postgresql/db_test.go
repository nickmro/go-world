package postgresql_test

import (
	"github.com/nickmro/world-app/world/db/postgresql"
	"github.com/drinkin/di/env"
)

// DB is a test wrapper for postgres.DB.
type DB struct {
	*postgresql.DB
}

// NewDB returns a new instance of DB.
func NewDB() *DB {
	db := &DB{DB: postgresql.NewDB(env.MustGet("DB_URL"))}
	return db
}

// MustOpenDB returns an new, open instance of DB.
func MustOpenDB() *DB {
	db := NewDB()
	if err := db.Open(); err != nil {
		panic(err)
	}
	return db
}

// MustBegin returns a new transaction.
func (db *DB) MustBegin() *Tx {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	return &Tx{tx}
}