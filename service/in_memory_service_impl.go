package service

import (
	"github.com/aahel/restapi/errors"
	"github.com/aahel/restapi/store"
	"github.com/aahel/restapi/types"
	"go.uber.org/zap"
)

type InMemoryServiceImpl struct {
	l     *zap.SugaredLogger
	store store.InMemoryStore
}

func NewInMemoryService(l *zap.SugaredLogger, store store.InMemoryStore) *InMemoryServiceImpl {
	return &InMemoryServiceImpl{l, store}
}

func (ims *InMemoryServiceImpl) PutData(kv *types.KeyValue) *types.KeyValue {
	keyVal := ims.store.PutData(kv)
	return keyVal
}

func (ims *InMemoryServiceImpl) GetData(key string) (*types.KeyValue, *errors.AppError) {
	kv, err := ims.store.GetData(key)
	if err != nil {
		return nil, err
	}
	return kv, err
}
