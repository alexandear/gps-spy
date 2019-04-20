package handler

import (
	"time"

	"github.com/pkg/errors"
)

func ParseTimestamp(ts string) (time.Time, error) {
	t, err := time.Parse("2006/01/02-15:04:05", ts)
	if err != nil {
		return time.Time{}, errors.New("timestamp must be in format 'YYYY/MM/DD-hh:mm:ss'")
	}
	loc, err := time.LoadLocation("Europe/Kiev")
	if err != nil {
		return time.Time{}, errors.Wrap(err, "failed to load location")
	}
	return t.In(loc), nil
}
