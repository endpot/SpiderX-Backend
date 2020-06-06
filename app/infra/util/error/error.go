package error

import (
	"encoding/json"
	"github.com/endpot/SpiderX-Backend/app/infra/util/http"
)

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}

func NewBadRequestError(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewUnauthorizedError(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func NewForbiddenError(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusForbidden,
		Message: message,
	}
}

func NewNotFoundError(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewInternalServerError(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
