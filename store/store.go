package store

import (
	"time"

	"github.com/aahel/restapi/entity"
	"github.com/aahel/restapi/errors"
	"github.com/aahel/restapi/types"
)

type RecordStore interface {
	GetRecords(tartDate, endDate time.Time, minCount int64, maxCount int64) ([]*entity.Record, *errors.AppError)
}

type InMemoryStore interface {
	PutData(kv *types.KeyValue) *types.KeyValue
	GetData(key string) (*types.KeyValue, *errors.AppError)
}
