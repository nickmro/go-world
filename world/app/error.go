package app

type Error string

const (
	ErrInternal = Error("internal_error")
)

func (e Error) Error() string {
	return string(e)
}
