package errs

import (
	"net/http"
)

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func NewNotFoundError(meessage string) error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: meessage,
	}
}

func NewUnexpectedError() error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "unexpected error",
	}
}

func NewValidatorError(meessage string) error {
	return AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: meessage,
	}
}
