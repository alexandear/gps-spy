package together

import (
	"time"

	"github.com/pkg/errors"

	"github.com/devchallenge/spy-api/internal/model"
	"github.com/devchallenge/spy-api/internal/util"
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
	if distance <= 0 {
		return 0, errors.Wrap(model.ErrInvalidArgument, "distance must be greater zero")
	}
	if number1 == number2 {
		return 0, errors.Wrap(model.ErrInvalidArgument, "numbers must not be equal")
	}

	items1, err := s.storage.Read(number1)
	if err != nil {
		return 0, errors.Wrap(err, "failed to read from storage for number1")
	}
	items2, err := s.storage.Read(number2)
	if err != nil {
		return 0, errors.Wrap(err, "failed to read from storage for number2")
	}

	items1 = excludeWorkingHours(items1)
	items2 = excludeWorkingHours(items2)

	items1 = excludeBeforeFrom(items1, from)
	items1 = excludeAfterTo(items1, to)
	items2 = excludeBeforeFrom(items2, from)
	items2 = excludeAfterTo(items2, to)

	var distances []float64
	for _, item := range items1 {
		distances = append(distances, distancesBetween(items2, item)...)
	}

	countAll := (len(items1) + len(items2)) / 2
	countTogether := 0
	for _, d := range distances {
		if d < float64(distance) {
			countTogether++
		}
	}
	if countTogether == 0 {
		return 0, nil
	}

	return (countAll / countTogether) * 100, nil
}

func excludeBeforeFrom(items []model.Together, from time.Time) []model.Together {
	res := make([]model.Together, 0, len(items))
	for _, item := range items {
		if !item.Timestamp.Before(from) {
			res = append(res, item)
		}
	}
	return res
}

func excludeAfterTo(items []model.Together, to time.Time) []model.Together {
	res := make([]model.Together, 0, len(items))
	for _, item := range items {
		if !item.Timestamp.After(to) {
			res = append(res, item)
		}
	}
	return res
}

func excludeWorkingHours(items []model.Together) []model.Together {
	const (
		workingHourStart = 9
		workingHourEnd   = 18
	)

	res := make([]model.Together, 0, len(items))
	for _, item := range items {
		hour := item.Timestamp.Hour()
		if hour < workingHourStart || hour > workingHourEnd {
			res = append(res, item)
		}
	}
	return res
}

func distancesBetween(from []model.Together, to model.Together) []float64 {
	res := make([]float64, 0, len(from))
	for _, item := range from {
		res = append(res, util.Distance(item.Coordinate, to.Coordinate))
	}
	return res
}
