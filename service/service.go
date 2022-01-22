package service

import (
	"time"

	"github.com/aahel/restapi/errors"
	"github.com/aahel/restapi/types"
)

type RecordService interface {
	GetRecords(startDate, endDate time.Time, minCount int64, maxCount int64) (*types.RecordResp, *errors.AppError)
}

type InMemoryService interface {
	GetData(key string) (*types.KeyValue, *errors.AppError)
	PutData(kv *types.KeyValue) *types.KeyValue
}
