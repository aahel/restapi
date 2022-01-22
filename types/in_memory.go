package types

import (
	"strings"

	"github.com/aahel/restapi/errors"
)

type KeyValue struct {
	// required: true
	// example: active-tabs
	Key string `json:"key,omitempty"`
	// required: true
	// example: getir
	Value string `json:"value,omitempty"`
}

func (kv *KeyValue) Ok() *errors.AppError {
	if strings.TrimSpace(kv.Key) == "" {
		return errors.KeyRequired()
	}
	if strings.TrimSpace(kv.Value) == "" {
		return errors.ValueRequired()
	}
	return nil
}
