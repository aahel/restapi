package store

import (
	"time"

	"github.com/aahel/restapi/errors"
	"github.com/aahel/restapi/model"
	"github.com/aahel/restapi/types"
)

type RecordStore interface {
	GetRecords(tartDate, endDate time.Time, minCount int64, maxCount int64) ([]*model.Record, *errors.AppError)
}

type InMemoryStore interface {
	PutData(kv *types.KeyValue) *types.KeyValue
	GetData(key string) (*types.KeyValue, *errors.AppError)
}
