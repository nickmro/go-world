package db

// DB is the interface for a DB integration.
type DB interface {
	Begin() (Tx, error)
}
