package db

const (
	ErrInternal = Error("internal_error")
)

const (
	ErrCountryNotFound = Error("country_not_found")
)

type Error string

func (e Error) Error() string {
	return string(e)
}
