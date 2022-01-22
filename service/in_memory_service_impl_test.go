package service

import (
	"testing"

	"github.com/aahel/restapi/config"
	"github.com/aahel/restapi/errors"
	"github.com/aahel/restapi/mocks"
	"github.com/aahel/restapi/types"
	"github.com/stretchr/testify/assert"
)

func TestPutData(t *testing.T) {
	mockInMemStore := new(mocks.InMemoryStore)
	kv := &types.KeyValue{
		Key:   "key",
		Value: "val",
	}
	mockInMemStore.On("PutData", kv).Return(kv)
	lgr := config.GetConsoleLogger()
	svc := NewInMemoryService(lgr, mockInMemStore)
	res := svc.PutData(kv)
	mockInMemStore.AssertExpectations(t)
	assert.NotNil(t, res)
	assert.Equal(t, kv, res)
}

func TestGetData(t *testing.T) {
	mockInMemStore := new(mocks.InMemoryStore)
	key := "key"
	expectedKv := &types.KeyValue{
		Key:   "key",
		Value: "val",
	}
	mockInMemStore.On("GetData", key).Return(expectedKv, nil)
	lgr := config.GetConsoleLogger()
	svc := NewInMemoryService(lgr, mockInMemStore)
	res, err := svc.GetData(key)
	mockInMemStore.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, expectedKv, res)
}

func TestGetDataKeyNotFound(t *testing.T) {
	mockInMemStore := new(mocks.InMemoryStore)
	key := "key"
	keyNotFoundErr := errors.KeyNotFound()
	mockInMemStore.On("GetData", key).Return(nil, keyNotFoundErr)
	lgr := config.GetConsoleLogger()
	svc := NewInMemoryService(lgr, mockInMemStore)
	res, err := svc.GetData(key)
	mockInMemStore.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Nil(t, res)
	assert.Equal(t, keyNotFoundErr, err)
}
