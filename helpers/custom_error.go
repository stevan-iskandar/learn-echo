package helpers

type Error struct {
	message string
}

func CustomError(message string) *Error {
	return &Error{message}
}

// Error returns the error message for CustomError.
func (e *Error) Error() string {
	return e.message
}
