package types

import (
	"net/http"

	"github.com/aahel/restapi/errors"
	"github.com/aahel/restapi/respond"
)

type Handler func(w http.ResponseWriter, r *http.Request) *errors.AppError

func (f Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := f(w, r)

	if err != nil {
		respond.Fail(w, err)
	}
}
