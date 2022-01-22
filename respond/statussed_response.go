package respond

import (
	"log"
	"net/http"

	"github.com/aahel/restapi/errors"
)

// OK is a helper function used to send response data
// with StatusOK status code (200)
func OK(w http.ResponseWriter, data, meta interface{}) *errors.AppError {
	return SendResponse(w, http.StatusOK, data, nil)
}

// Created is a helper function used to send response data
// with StatusCreated status code (201)
func Created(w http.ResponseWriter, data, meta interface{}) *errors.AppError {
	return SendResponse(w, http.StatusCreated, data, nil)
}

// Fail write the error response
// Common func to send all the error response
func Fail(w http.ResponseWriter, e *errors.AppError) {
	log.Printf("Code: %d, Error: %s\n DEBUG: %s\n",
		e.Status, e.Error(), e.Debug)
	SendResponse(w, e.Status, e, nil)
}
