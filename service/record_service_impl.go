package service

import (
	"time"

	"github.com/aahel/restapi/errors"
	"github.com/aahel/restapi/store"
	"github.com/aahel/restapi/types"
	"go.uber.org/zap"
)

type RecordServiceImpl struct {
	l     *zap.SugaredLogger
	store store.RecordStore
}

func NewRecordService(l *zap.SugaredLogger, db store.RecordStore) *RecordServiceImpl {
	return &RecordServiceImpl{l, db}
}

func (rs *RecordServiceImpl) GetRecords(startDate, endDate time.Time, minCount int64, maxCount int64) (*types.RecordResp, *errors.AppError) {
	records, err := rs.store.GetRecords(startDate, endDate, minCount, maxCount)
	if err != nil {
		return nil, err
	}
	recordResp := &types.RecordResp{
		Code:    0,
		Msg:     "Success",
		Records: records,
	}
	return recordResp, nil
}
