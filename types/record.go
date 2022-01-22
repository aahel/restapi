package types

import (
	"strings"
	"time"

	"github.com/aahel/restapi/entity"
	"github.com/aahel/restapi/errors"
)

type RecordFilterReq struct {
	// required: true
	// example: "2016-05-24"
	StartDate string `json:"startDate,omitempty"`
	// required: true
	// example: "2016-07-24"
	EndDate string `json:"endDate,omitempty"`
	// required: true
	// example: 2700
	MinCount int64 `json:"minCount,omitempty"`
	// required: true
	//example:3700
	MaxCount int64 `json:"maxCount,omitempty"`
}

func (rr *RecordFilterReq) Ok() *errors.AppError {
	if strings.TrimSpace(rr.StartDate) == "" || strings.TrimSpace(rr.EndDate) == "" {
		return errors.InavalidDates()
	}
	if rr.MinCount > rr.MaxCount {
		return errors.InavalidCountFilter()
	}
	return nil
}

type RecordResp struct {
	// example: 0
	Code int `json:"code"`
	// example: "Success"
	Msg     string           `json:"msg,omitempty"`
	Records []*entity.Record `json:"records,omitempty"`
}

type Record struct {
	// example: "fhhggjkb"
	Key string `json:"key,omitempty"`
	// example: "2016-12-28T09:25:32.818Z"
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// example: 2800
	TotalCount int64 `json:"totalCount,omitempty"`
}
