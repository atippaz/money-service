package utils

import "time"

type Datetime interface {
	Now() (time.Time, error)
	GetStartDate(time ...time.Time) (time.Time, error)
	GetEndDate(time ...time.Time) (time.Time, error)
	ConvertToDate(time string) (time.Time, error)
}
type datetime struct {
}
