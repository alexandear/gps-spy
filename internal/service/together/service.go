package together

import (
	"math/rand"
	"time"

	"github.com/pkg/errors"

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

func (s *Service) SpendPercentage(number1, number2 string, from, to time.Time, distance int) (int, error) {
	if to.Before(from) {
		return 0, errors.Wrap(model.ErrInvalidArgument, "to must be greater from")
	}
	return rand.Intn(101), nil
}
