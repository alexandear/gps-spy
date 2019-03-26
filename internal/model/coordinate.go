package model

import "github.com/pkg/errors"

type Coordinate struct {
	longitude float32
	latitude  float32
}

func NewCoordinate(longitude, latitude float32) (*Coordinate, error) {
	if longitude < -180 || longitude > 180 {
		return nil, errors.New("longitude must be in range [-180; 180]")
	}
	if latitude < -90 || latitude > 90 {
		return nil, errors.New("latitude must be in range [-90; 90]")
	}
	return &Coordinate{
		longitude: longitude,
		latitude:  latitude,
	}, nil
}
