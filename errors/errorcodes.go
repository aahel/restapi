package errors

const (
	InvalidDateCode    = 1
	InvalidDateMessage = "date is invalid"

	InvalidCountFilterCode    = 2
	InvalidCountFilterMessage = "count filter is invalid"

	UmarshallErrorCode    = 3
	UmarshallErrorMessage = "failed to unmarshall json"

	KeyNotFoundCode    = 4
	KeyNotFoundMessage = "key not found"

	KeyRequiredCode    = 5
	KeyRequiredMessage = "key required"

	ValRequiredCode    = 6
	ValRequiredMessage = "value required"

	InternalServerStdCode    = 7
	InternalServerStdMessage = "Something went wrong"

	RecordNotFoundCode    = 8
	RecordNotFoundMessage = "record not found"
)
