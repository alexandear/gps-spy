package handler

import (
	"time"

	"github.com/devchallenge/spy-api/internal/gen/restapi/operations"
	"github.com/devchallenge/spy-api/internal/model"
)

type Handler struct {
	gps      GPS
	together Together
	violator Violator
}

type GPS interface {
	Add(phone model.Phone, coordinate model.Coordinate, timestamp time.Time) error
}

type Together interface {
	SpendPercentage(number1, number2 string, from, to time.Time, distance int) (int, error)
}

type Violator interface {
	Numbers(from, to time.Time, minSpeed int, minCoordinate, maxCoordinate model.Coordinate) ([]string, error)
}

func New(gps GPS, together Together, violator Violator) *Handler {
	return &Handler{
		gps:      gps,
		together: together,
		violator: violator,
	}
}

func (h *Handler) ConfigureHandlers(api *operations.SpyAPI) {
	api.PostBbinputHandler = operations.PostBbinputHandlerFunc(h.PostBbinputHandler)
	api.PostBbsHandler = operations.PostBbsHandlerFunc(h.PostBbsHandler)
	api.PostBbfastDriveHandler = operations.PostBbfastDriveHandlerFunc(h.PostBbfastDriveHandler)
}
