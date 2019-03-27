package model

import (
	"github.com/pkg/errors"
)

type Coordinate struct {
	Longitude float32
	Latitude  float32
}

func NewCoordinate(longitude, latitude float32) (Coordinate, error) {
	if longitude < -180 || longitude > 180 {
		return Coordinate{}, errors.Wrap(ErrInvalidArgument, "longitude must be in range [-180; 180]")
	}
	if latitude < -90 || latitude > 90 {
		return Coordinate{}, errors.Wrap(ErrInvalidArgument, "latitude must be in range [-90; 90]")
	}
	return Coordinate{
		Longitude: longitude,
		Latitude:  latitude,
	}, nil
}
