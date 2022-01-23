package v1

import (
	"bytes"
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aahel/restapi/config"
	"github.com/aahel/restapi/errors"
	"github.com/aahel/restapi/mocks"
	"github.com/aahel/restapi/model"
	"github.com/aahel/restapi/types"
	"github.com/aahel/restapi/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetRecords(t *testing.T) {
	mockService := new(mocks.RecordService)
	recordReq := &types.RecordFilterReq{
		StartDate: "2016-01-26",
		EndDate:   "2016-05-27",
		MinCount:  2700,
		MaxCount:  3700,
	}
	jsonBytes, errz := json.Marshal(recordReq)
	assert.Nil(t, errz)
	r := httptest.NewRequest("POST", "localhost:8081/v1/records", bytes.NewBuffer(jsonBytes))
	rec := httptest.NewRecorder()
	startTime, _ := utils.StrToTime("2016-01-26")
	endTime, _ := utils.StrToTime("2016-05-27")
	minCount := int64(2700)
	maxCount := int64(3700)
	recordRes := &types.RecordResp{
		Code:    0,
		Msg:     "Success",
		Records: []*model.Record{{Key: "ihfhahf", TotalCount: 2900, CreatedAt: startTime.AddDate(0, 0, 4)}},
	}
	mockService.On("GetRecords", startTime, endTime, minCount, maxCount).Return(recordRes, nil)
	lgr := config.GetConsoleLogger()
	recHandl := NewRecordHandler(lgr, mockService)
	err := recHandl.GetRecords(rec, r)
	mockService.AssertExpectations(t)
	assert.Nil(t, err)
	res := rec.Result()
	defer res.Body.Close()
	errx := json.NewDecoder(res.Body).Decode(recordRes)
	assert.Nil(t, errx)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
}

func TestGetRecordsMinCountInvalid(t *testing.T) {
	mockService := new(mocks.RecordService)
	recordReq := &types.RecordFilterReq{
		StartDate: "2016-01-26",
		EndDate:   "2016-05-27",
		MinCount:  2700,
		MaxCount:  2000,
	}
	jsonBytes, errz := json.Marshal(recordReq)
	assert.Nil(t, errz)
	r := httptest.NewRequest("POST", "localhost:8081/v1/records", bytes.NewBuffer(jsonBytes))
	rec := httptest.NewRecorder()

	lgr := config.GetConsoleLogger()
	recHandl := NewRecordHandler(lgr, mockService)
	err := recHandl.GetRecords(rec, r)
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.InavalidCountFilter())
}
func TestGetRecordsInvalidStartDate(t *testing.T) {
	mockService := new(mocks.RecordService)
	recordReq := &types.RecordFilterReq{
		StartDate: "2016-01",
		EndDate:   "2016-05-27",
		MinCount:  2700,
		MaxCount:  3700,
	}
	jsonBytes, errz := json.Marshal(recordReq)
	assert.Nil(t, errz)
	r := httptest.NewRequest("POST", "localhost:8081/v1/records", bytes.NewBuffer(jsonBytes))
	rec := httptest.NewRecorder()

	lgr := config.GetConsoleLogger()
	recHandl := NewRecordHandler(lgr, mockService)
	err := recHandl.GetRecords(rec, r)
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.InavalidDates())
}

func TestGetRecordsInvalidEndDate(t *testing.T) {
	mockService := new(mocks.RecordService)
	recordReq := &types.RecordFilterReq{
		StartDate: "2016-01-26",
		EndDate:   "2016-05",
		MinCount:  2700,
		MaxCount:  3700,
	}
	jsonBytes, errz := json.Marshal(recordReq)
	assert.Nil(t, errz)
	r := httptest.NewRequest("POST", "localhost:8081/v1/records", bytes.NewBuffer(jsonBytes))
	rec := httptest.NewRecorder()

	lgr := config.GetConsoleLogger()
	recHandl := NewRecordHandler(lgr, mockService)
	err := recHandl.GetRecords(rec, r)
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.InavalidDates())
}

func TestGetRecordsEndDateBeforeStatDateErr(t *testing.T) {
	mockService := new(mocks.RecordService)
	recordReq := &types.RecordFilterReq{
		StartDate: "2016-01-26",
		EndDate:   "2015-05-27",
		MinCount:  2700,
		MaxCount:  3700,
	}
	jsonBytes, errz := json.Marshal(recordReq)
	assert.Nil(t, errz)
	r := httptest.NewRequest("POST", "localhost:8081/v1/records", bytes.NewBuffer(jsonBytes))
	rec := httptest.NewRecorder()

	lgr := config.GetConsoleLogger()
	recHandl := NewRecordHandler(lgr, mockService)
	err := recHandl.GetRecords(rec, r)
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.InvalidDateRange())
}

func TestGetRecordsNotFoundErr(t *testing.T) {
	mockService := new(mocks.RecordService)
	recordReq := &types.RecordFilterReq{
		StartDate: "2016-01-26",
		EndDate:   "2016-05-27",
		MinCount:  2700,
		MaxCount:  3700,
	}
	jsonBytes, errz := json.Marshal(recordReq)
	assert.Nil(t, errz)
	r := httptest.NewRequest("POST", "localhost:8081/v1/records", bytes.NewBuffer(jsonBytes))
	rec := httptest.NewRecorder()
	startTime, _ := utils.StrToTime("2016-01-26")
	endTime, _ := utils.StrToTime("2016-05-27")
	minCount := int64(2700)
	maxCount := int64(3700)
	errs := errors.RecordNotFound()
	mockService.On("GetRecords", startTime, endTime, minCount, maxCount).Return(nil, errs)
	lgr := config.GetConsoleLogger()
	recHandl := NewRecordHandler(lgr, mockService)
	err := recHandl.GetRecords(rec, r)
	mockService.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Equal(t, errs, err)
}
