package postgresql

import (
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
	"gopkg.in/mgutz/dat.v1"
	"github.com/huandu/xstrings"
	"github.com/jmoiron/sqlx"
	"github.com/drinkin/di/lg"
	"github.com/nickmro/world-app/world/db"
)

var _ db.DB = &DB{}

// DB wraps a postgresql.DB instance.
type DB struct {
	URL string
	runner *runner.DB
	logger *lg.Logger
}

// NewDB returns a new database instance.
func NewDB(url string) *DB {
	return &DB{
		URL: url,
		logger: lg.New("postgresql"),
	}
}

// Open opens and initializes the PostgreSQL database.
func (db *DB) Open() error {
	dbx, err := sqlx.Connect("postgres", db.URL)
	if err != nil {
		db.logger.Info(err.Error())
		db.logger.Error(err.Error())
		return err
	}

	dbx.MapperFunc(xstrings.ToSnakeCase)
	dat.EnableInterpolation = true

	db.runner = runner.NewDBFromSqlx(dbx)

	return nil
}

// Close closes the connection to the database.
func (db *DB) Close() error {
	if db.runner != nil {
		return db.runner.DB.Close()
	}
	return nil
}

// Begin begins a transaction.
func (db *DB) Begin() (db.Tx, error) {
	tx, err := db.runner.Begin()
	if err != nil {
		db.logger.Info(err.Error())
		db.logger.Error(err.Error())
		return nil, err
	}
	return &Tx{tx, db.logger}, nil
}


