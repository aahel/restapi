package utils

import "time"

const (
	EMPTY_STRING = ""
	time_layout  = "2006-01-02"
)

// Converts string date to time
func StrToTime(date string) (time.Time, error) {
	return time.Parse(time_layout, date)
}
