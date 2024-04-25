package customError

import "net/http"

type CustomError struct {
	Err        error
	StatusCode int
}

func NewInternalError(err error) *CustomError {
	return &CustomError{
		Err:        err,
		StatusCode: http.StatusInternalServerError,
	}
}

func (e *CustomError) Error() string {
	return e.Err.Error()
}

func NewBadRequestError(err error) *CustomError {
	return &CustomError{
		Err:        err,
		StatusCode: http.StatusBadRequest,
	}
}
