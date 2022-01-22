package v1

import (
	"net/http"
	"strings"

	"github.com/aahel/restapi/errors"
	"github.com/aahel/restapi/respond"
	"github.com/aahel/restapi/service"
	"github.com/aahel/restapi/types"
	"github.com/aahel/restapi/types/consts"
	"github.com/aahel/restapi/utils"
	"go.uber.org/zap"
)

type InMemoryHandler struct {
	lgr *zap.SugaredLogger
	svc service.InMemoryService
}

func NewInMemoryHandler(lgr *zap.SugaredLogger, svc service.InMemoryService) *InMemoryHandler {
	return &InMemoryHandler{lgr, svc}
}

// swagger:route POST /v1/in-memory inMemory inMemoryReq
// stores key value in in memory db
// responses:
// 200: inMemorySuccessResponse
// 400: errResp
func (im *InMemoryHandler) WriteData(rw http.ResponseWriter, r *http.Request) *errors.AppError {
	var kvReq types.KeyValue
	if err := utils.Decode(r, &kvReq); err != nil {
		return err
	}
	kv := im.svc.PutData(&kvReq)
	return respond.OK(rw, kv, nil)
}

// swagger:route GET /v1/in-memory inMemory inMemoryQueryParams
// stores key value in in memory db
// responses:
// 200: inMemorySuccessResponse
// 400: errResp
// 404: errResp
func (im *InMemoryHandler) GetData(rw http.ResponseWriter, r *http.Request) *errors.AppError {
	key := r.URL.Query().Get(consts.KEY)
	if strings.TrimSpace(key) == "" {
		return errors.KeyRequired()
	}
	kv, err := im.svc.GetData(key)
	if err != nil {
		return err
	}
	return respond.OK(rw, kv, nil)
}
