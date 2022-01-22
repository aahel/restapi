package service

import (
	"testing"

	"github.com/aahel/restapi/config"
	"github.com/aahel/restapi/entity"
	"github.com/aahel/restapi/errors"
	"github.com/aahel/restapi/mocks"
	"github.com/aahel/restapi/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetRecords(t *testing.T) {
	mockRecordStore := new(mocks.RecordStore)
	lgr := config.GetConsoleLogger()
	svc := NewRecordService(lgr, mockRecordStore)
	startTime, _ := utils.StrToTime("2016-01-26")
	endTime, _ := utils.StrToTime("2016-05-27")
	minCount := int64(2700)
	maxCount := int64(3700)
	expectedRecords := []*entity.Record{{Key: "ihfhahf", TotalCount: 2900, CreatedAt: startTime.AddDate(0, 0, 4)}}
	mockRecordStore.On("GetRecords", startTime, endTime, minCount, maxCount).Return(expectedRecords, nil)
	recs, err := svc.GetRecords(startTime, endTime, minCount, maxCount)
	mockRecordStore.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, recs)
	assert.Equal(t, recs.Code, 0)
	assert.Equal(t, recs.Msg, "Success")
	assert.Equal(t, recs.Records, expectedRecords)
}

func TestGetRecordsNotFound(t *testing.T) {
	mockRecordStore := new(mocks.RecordStore)
	lgr := config.GetConsoleLogger()
	svc := NewRecordService(lgr, mockRecordStore)
	startTime, _ := utils.StrToTime("2016-01-26")
	endTime, _ := utils.StrToTime("2016-05-27")
	minCount := int64(2700)
	maxCount := int64(3700)
	recordNotFoundErr := errors.RecordNotFound()
	mockRecordStore.On("GetRecords", startTime, endTime, minCount, maxCount).Return(nil, recordNotFoundErr)
	recs, err := svc.GetRecords(startTime, endTime, minCount, maxCount)
	mockRecordStore.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Nil(t, recs)
	assert.Equal(t, recordNotFoundErr, err)
}
