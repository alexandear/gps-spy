package storage

import (
	"github.com/pkg/errors"

	"github.com/devchallenge/spy-api/internal/model"
)

type Storage struct {
}

func (s *Storage) SaveLocation(location *model.Location) error {
	if location == nil {
		return errors.New("location must be not empty")
	}
	return nil
}
