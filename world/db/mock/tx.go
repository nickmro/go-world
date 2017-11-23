package mock

import (
	"github.com/nickmro/world-app/world"
)

type Tx struct {
	GetCountryByCodeFn func(string) (*world.Country, error)
}

// NewTx returns a default mock Tx.
func NewTx() *Tx {
	getCountryByCodeFn := func(string) (*world.Country, error) {
		return &world.Country{}, nil
	}

	tx := &Tx{
		GetCountryByCodeFn: getCountryByCodeFn,
	}

	return tx
}

func (tx *Tx) Commit() error {
	return nil
}

func (tx *Tx) Rollback() error {
	return nil
}

func (tx *Tx) AutoRollback() error {
	return nil
}
