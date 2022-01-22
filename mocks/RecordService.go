// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	errors "github.com/aahel/restapi/errors"
	mock "github.com/stretchr/testify/mock"

	time "time"

	types "github.com/aahel/restapi/types"
)

// RecordService is an autogenerated mock type for the RecordService type
type RecordService struct {
	mock.Mock
}

// GetRecords provides a mock function with given fields: startDate, endDate, minCount, maxCount
func (_m *RecordService) GetRecords(startDate time.Time, endDate time.Time, minCount int64, maxCount int64) (*types.RecordResp, *errors.AppError) {
	ret := _m.Called(startDate, endDate, minCount, maxCount)

	var r0 *types.RecordResp
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, int64, int64) *types.RecordResp); ok {
		r0 = rf(startDate, endDate, minCount, maxCount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.RecordResp)
		}
	}

	var r1 *errors.AppError
	if rf, ok := ret.Get(1).(func(time.Time, time.Time, int64, int64) *errors.AppError); ok {
		r1 = rf(startDate, endDate, minCount, maxCount)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.AppError)
		}
	}

	return r0, r1
}
