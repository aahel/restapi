// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	entity "github.com/aahel/restapi/entity"
	errors "github.com/aahel/restapi/errors"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// RecordStore is an autogenerated mock type for the RecordStore type
type RecordStore struct {
	mock.Mock
}

// GetRecords provides a mock function with given fields: tartDate, endDate, minCount, maxCount
func (_m *RecordStore) GetRecords(tartDate time.Time, endDate time.Time, minCount int64, maxCount int64) ([]*entity.Record, *errors.AppError) {
	ret := _m.Called(tartDate, endDate, minCount, maxCount)

	var r0 []*entity.Record
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, int64, int64) []*entity.Record); ok {
		r0 = rf(tartDate, endDate, minCount, maxCount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Record)
		}
	}

	var r1 *errors.AppError
	if rf, ok := ret.Get(1).(func(time.Time, time.Time, int64, int64) *errors.AppError); ok {
		r1 = rf(tartDate, endDate, minCount, maxCount)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.AppError)
		}
	}

	return r0, r1
}