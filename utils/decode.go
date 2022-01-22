package utils

import (
	"encoding/json"
	"net/http"

	"github.com/aahel/restapi/errors"
)

type ok interface {
	Ok() *errors.AppError
}

// Decode - decodes the request body and extends the validator interface
func Decode(r *http.Request, v interface{}) *errors.AppError {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return errors.UmarshallError()
	}

	if payload, ok := v.(ok); ok {
		return payload.Ok()
	}
	return nil
}
