package store

import (
	"sync"

	"github.com/aahel/restapi/errors"
	"github.com/aahel/restapi/types"
	"go.uber.org/zap"
)

type InMemoryStoreImpl struct {
	l         *zap.SugaredLogger
	datastore map[string]string
	lock      sync.RWMutex
}

func NewInMemoryStore(l *zap.SugaredLogger, datastore map[string]string) *InMemoryStoreImpl {
	return &InMemoryStoreImpl{l: l, datastore: datastore}
}

func (ims *InMemoryStoreImpl) PutData(kv *types.KeyValue) *types.KeyValue {
	ims.lock.Lock()
	defer ims.lock.Unlock()
	ims.datastore[kv.Key] = kv.Value
	return kv
}

func (ims *InMemoryStoreImpl) GetData(key string) (*types.KeyValue, *errors.AppError) {
	ims.lock.RLock()
	defer ims.lock.RUnlock()
	value := ims.datastore[key]
	if value == "" {
		return nil, errors.KeyNotFound()
	}
	return &types.KeyValue{Key: key, Value: ims.datastore[key]}, nil
}
