package model

import (
	"net"
	"time"

	"github.com/pkg/errors"
)

type Location struct {
	number     string
	imei       string
	ip         net.IP
	timestamp  time.Time
	coordinate Coordinate
}

func NewLocation(number, imei string, longitude, latitude float32) (*Location, error) {
	if number == "" {
		return nil, errors.New("number must be not empty")
	}
	if imei == "" {
		return nil, errors.New("number must be not empty")
	}
	coordinate, err := NewCoordinate(longitude, latitude)
	if err != nil {
		return nil, errors.Wrap(err, "coordinate must be valid")
	}
	return &Location{
		number:     number,
		imei:       imei,
		coordinate: *coordinate,
	}, nil
}

func (l *Location) SetIP(ip net.IP) {
	l.ip = ip
}

func (l *Location) SetTimestamp(timestamp time.Time) {
	l.timestamp = timestamp
}
