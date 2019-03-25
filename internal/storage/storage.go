package storage

import (
	"github.com/devchallenge/spy-api/internal/service/model"
	"github.com/pkg/errors"
)

type Storage struct {
}

func (s *Storage) SaveLocation(location *model.Location) error {
	if location == nil {
		return errors.New("location must be not empty")
	}
	return nil
}
