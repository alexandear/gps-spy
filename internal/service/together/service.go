package together

import (
	"time"

	"github.com/devchallenge/spy-api/internal/model"
)

type Storage interface {
	Read(number string) ([]model.Together, error)
}

type Service struct {
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) SpendPercentage(number1, number2 string, from, to time.Time) (int, error) {
	// TODO
	return 0, nil
}
