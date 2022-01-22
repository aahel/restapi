package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aahel/restapi/config"
	"github.com/aahel/restapi/errors"
	"github.com/aahel/restapi/mocks"
	"github.com/aahel/restapi/types"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestWriteData(t *testing.T) {
	mockService := new(mocks.InMemoryService)
	kv := &types.KeyValue{
		Key:   "key",
		Value: "val",
	}
	jsonBytes, errz := json.Marshal(kv)
	assert.Nil(t, errz)
	r := httptest.NewRequest("POST", "localhost:8081/v1/in-memory", bytes.NewBuffer(jsonBytes))
	rec := httptest.NewRecorder()
	mockService.On("PutData", kv).Return(kv, nil)
	lgr := config.GetConsoleLogger()
	imHandl := NewInMemoryHandler(lgr, mockService)
	err := imHandl.WriteData(rec, r)
	mockService.AssertExpectations(t)
	assert.Nil(t, err)
	res := rec.Result()
	defer res.Body.Close()
	errx := json.NewDecoder(res.Body).Decode(kv)
	assert.Nil(t, errx)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
}

func TestWriteDataValueRequired(t *testing.T) {
	mockService := new(mocks.InMemoryService)
	kv := &types.KeyValue{
		Key: "key",
	}
	jsonBytes, errz := json.Marshal(kv)
	assert.Nil(t, errz)
	r := httptest.NewRequest("POST", "localhost:8081/v1/in-memory", bytes.NewBuffer(jsonBytes))
	rec := httptest.NewRecorder()
	lgr := config.GetConsoleLogger()
	imHandl := NewInMemoryHandler(lgr, mockService)
	err := imHandl.WriteData(rec, r)
	assert.NotNil(t, err)
	assert.Equal(t, errors.ValueRequired(), err)
}

func TestGetData(t *testing.T) {
	key := "test"
	r := httptest.NewRequest("GET", "localhost:8080/v1/in-memory?key="+key, nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("key", key)
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	rec := httptest.NewRecorder()
	mockService := new(mocks.InMemoryService)
	kv := &types.KeyValue{
		Key:   "key",
		Value: "val",
	}
	mockService.On("GetData", key).Return(kv, nil)
	lgr := config.GetConsoleLogger()
	imHandl := NewInMemoryHandler(lgr, mockService)
	err := imHandl.GetData(rec, r)
	mockService.AssertExpectations(t)
	assert.Nil(t, err)
	res := rec.Result()
	defer res.Body.Close()
	errx := json.NewDecoder(res.Body).Decode(kv)
	assert.Nil(t, errx)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
}

func TestGetDatakeyRequired(t *testing.T) {
	key := "test"
	r := httptest.NewRequest("GET", "localhost:8080/v1/in-memory?key=", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("key", key)
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	rec := httptest.NewRecorder()
	mockService := new(mocks.InMemoryService)
	lgr := config.GetConsoleLogger()
	imHandl := NewInMemoryHandler(lgr, mockService)
	err := imHandl.GetData(rec, r)
	assert.NotNil(t, err)
	assert.Equal(t, errors.KeyRequired(), err)
}

func TestGetDataKeynotFound(t *testing.T) {
	key := "test"
	r := httptest.NewRequest("GET", "localhost:8080/v1/in-memory?key="+key, nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("key", key)
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	rec := httptest.NewRecorder()
	mockService := new(mocks.InMemoryService)
	errs := errors.KeyNotFound()
	mockService.On("GetData", key).Return(nil, errs)
	lgr := config.GetConsoleLogger()
	imHandl := NewInMemoryHandler(lgr, mockService)
	err := imHandl.GetData(rec, r)
	mockService.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Equal(t, errs, err)
}
