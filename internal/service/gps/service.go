package gps

import (
	"time"

	"github.com/pkg/errors"

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
	if err := s.storage.Save(phone, coordinate, timestamp); err != nil {
		return errors.Wrap(err, "failed to save")
	}
	return nil
}
