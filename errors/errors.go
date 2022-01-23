package errors

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

// swagger:response errResp
// AppError struct holds the value of HTTP status code and custom error message.
type AppError struct {
	//in: body

	Status  int    `json:"-"`
	Code    int    `json:"code"`
	Message string `json:"msg,omitempty"`
	Debug   error  `json:"-"`
}

func (err *AppError) Error() string {
	return err.Message
}

// AddDebug method is used to add a debug error which will be printed
// during the error execution if it is not nil. This is purely for developers'
// debugging purposes
func (err *AppError) AddDebug(erx error) *AppError {
	if err != nil {
		err.Debug = erx
	}

	return err
}

// NewAppError returns the new apperror object
func NewAppError(status, code int, message string) *AppError {
	return &AppError{
		Status:  status,
		Code:    code,
		Message: message,
	}
}

func InavalidDates() *AppError {
	return NewAppError(http.StatusBadRequest, InvalidDateCode, InvalidDateMessage)
}

func InavalidCountFilter() *AppError {
	return NewAppError(http.StatusBadRequest, InvalidCountFilterCode, InvalidCountFilterMessage)
}

func KeyRequired() *AppError {
	return NewAppError(http.StatusBadRequest, KeyRequiredCode, KeyRequiredMessage)
}

func ValueRequired() *AppError {
	return NewAppError(http.StatusBadRequest, ValRequiredCode, ValRequiredMessage)
}

// NotFound will return code 2 with custom message.
func KeyNotFound() *AppError {
	return NewAppError(http.StatusNotFound, KeyNotFoundCode, KeyNotFoundMessage)
}

// InternalServerStd will return code 3 with static message.
func InternalServerStd() *AppError {
	return NewAppError(http.StatusInternalServerError, InternalServerStdCode, InternalServerStdMessage)
}

// IsMongoNoDocErr should return true if the err is redis: nil
func IsMongoNoDocErr(err error) bool {
	return err == mongo.ErrNoDocuments
}

func UmarshallError() *AppError {
	return NewAppError(http.StatusBadRequest, UmarshallErrorCode, UmarshallErrorMessage)
}

func RecordNotFound() *AppError {
	return NewAppError(http.StatusNotFound, RecordNotFoundCode, RecordNotFoundMessage)
}

func InvalidDateRange() *AppError {
	return NewAppError(http.StatusBadRequest, InvalidDateRangeCode, InvalidDateRangeMessage)
}
