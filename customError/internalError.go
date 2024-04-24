package customError

import "net/http"

type CustomError struct {
	Err        error
	StatusCode int
}

type InternalError struct {
	*CustomError
}

func NewInternalError(err error) *InternalError {
	return &InternalError{
		CustomError: &CustomError{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		},
	}
}

func (e *CustomError) Error() string {
	return e.Err.Error()
}

type BadRequestError struct {
	*CustomError
}

func NewBadRequestError(err error) *BadRequestError {
	return &BadRequestError{
		&CustomError{
			Err:        err,
			StatusCode: http.StatusBadRequest,
		},
	}
}
