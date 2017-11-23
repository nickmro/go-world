package db

// Tx is the interface for a database transaction integration.
type Tx interface {
	CountryTx

	Commitable
	Rollbackable
	AutoRollbackable
}

type Commitable interface {
	Commit() error
}

type Rollbackable interface {
	Rollback() error
}

type AutoRollbackable interface {
	AutoRollback() error
}
