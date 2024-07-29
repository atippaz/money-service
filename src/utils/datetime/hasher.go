package utils

import (
	"errors"
	"time"
)

func NewDatetime() Datetime {
	return &datetime{}
}
func (s *datetime) Now() (time.Time, error) {
	return time.Now().UTC(), nil
}
func (s *datetime) GetStartDate(t ...time.Time) (time.Time, error) {
	var date time.Time
	if len(t) > 0 {
		date = t[0]
	} else {
		date = time.Now().UTC()
	}
	if date.IsZero() {
		return time.Time{}, errors.New("invalid time provided")
	}
	year, month, day := date.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, date.Location()), nil
}
func (s *datetime) GetEndDate(t ...time.Time) (time.Time, error) {
	var date time.Time
	if len(t) > 0 {
		date = t[0]
	} else {
		date = time.Now().UTC()
	}
	if date.IsZero() {
		return time.Time{}, errors.New("invalid time provided")
	}
	startOfDay, err := s.GetStartDate(date)
	if err != nil {
		return time.Time{}, err
	}
	endOfDay := startOfDay.Add(24 * time.Hour).Add(-time.Second)
	return endOfDay, nil
}
func (s *datetime) ConvertToDate(_time string) (time.Time, error) {
	datetime, err := time.Parse(time.RFC3339Nano, _time)
	if err != nil {
		return time.Time{}, err
	}
	return datetime, nil
}
