package gps

import (
	"time"

	"github.com/devchallenge/spy-api/internal/model"
)

type Storage interface {
	Save(phone model.Phone, coordinate model.Coordinate, timestamp time.Time) error
}

type Service struct {
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) Add(phone model.Phone, coordinate model.Coordinate, timestamp time.Time) error {
	return nil
}
