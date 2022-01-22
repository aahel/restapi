package v1

import (
	"net/http"

	"github.com/aahel/restapi/errors"
	"github.com/aahel/restapi/respond"
	"github.com/aahel/restapi/service"
	"github.com/aahel/restapi/types"
	"github.com/aahel/restapi/utils"
	"go.uber.org/zap"
)

type RecordHandler struct {
	lgr *zap.SugaredLogger
	svc service.RecordService
}

func NewRecordHandler(lgr *zap.SugaredLogger, svc service.RecordService) *RecordHandler {
	return &RecordHandler{lgr, svc}
}

// swagger:route POST /v1/records records recordReq
// Return records from the database
// responses:
// 200: recordSuccessResponse
// 500: errResp
// 404: errResp
// 400: errResp
func (rh *RecordHandler) GetRecords(rw http.ResponseWriter, r *http.Request) *errors.AppError {
	var recordReq types.RecordFilterReq
	if err := utils.Decode(r, &recordReq); err != nil {
		return err
	}
	startDate, errx := utils.StrToTime(recordReq.StartDate)
	if errx != nil {
		return errors.InavalidDates()
	}
	endDate, erry := utils.StrToTime(recordReq.EndDate)
	if erry != nil {
		return errors.InavalidDates()
	}
	if startDate.After(endDate) {
		return errors.InavalidDates()
	}
	records, errz := rh.svc.GetRecords(startDate, endDate, recordReq.MinCount, recordReq.MaxCount)
	if errz != nil {
		return errz
	}
	return respond.OK(rw, records, nil)
}
