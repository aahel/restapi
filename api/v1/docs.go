package v1

import "github.com/aahel/restapi/types"

// swagger:parameters recordReq
type RecordsBody struct {
	//in: body
	//required: true
	Body types.RecordFilterReq
}

// swagger:response recordSuccessResponse
type RecordSuccessResponse struct {
	//in: body
	Body types.RecordResp
}

// swagger:parameters inMemoryReq
type InMemoryBody struct {
	//required: true
	//in: body
	Body types.KeyValue
}

// swagger:response inMemorySuccessResponse
type InMemorySuccessResponse struct {
	//in: body
	Body types.KeyValue
}

// swagger:parameters inMemoryQueryParams
type InMemoryQueryParams struct {
	//required: true
	Key string `json:"key,omitempty"`
}
