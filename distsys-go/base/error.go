package base

import (
	"fmt"
	"net/http"
)

const (
	CodeTokenNotTransmitted = 600
	CodeTokenNotParsed      = 601
	CodeTokenNotExisted     = 602
	CodeTokenExpired        = 603
	//User Model Error Code start from 7xx
	CodeUserNotCreated           = 700
	CodeUserIsForbidden          = 701
	CodeUserNotActiveOrIsDeleted = 702
)

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("CustomError, code: %d, message: %s", e.Code, e.Message)
}

func BuildCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func ConvertToCustomError(code int, err error) *CustomError {
	return &CustomError{
		Code:    code,
		Message: err.Error(),
	}
}

func InternalCustomError(err error) *CustomError {
	return &CustomError{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	}
}

func BuildCustomInternalError(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func BadRequestCustomError(err error) *CustomError {
	return &CustomError{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}
}

func BuildCustomBadReqError(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func ForbiddenCustomError(err error) *CustomError {
	return &CustomError{
		Code:    http.StatusForbidden,
		Message: err.Error(),
	}
}

func BuildCustomForbiddenError(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusForbidden,
		Message: message,
	}
}
