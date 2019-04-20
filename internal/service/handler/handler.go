package handler

import (
	"time"

	"github.com/devchallenge/spy-api/internal/gen/restapi/operations"
	"github.com/devchallenge/spy-api/internal/model"
)

type Handler struct {
	gps      GPS
	together Together
}

type GPS interface {
	Add(phone model.Phone, coordinate model.Coordinate, timestamp time.Time) error
}

type Together interface {
	SpendPercentage(number1, number2 string, from, to time.Time) (int, error)
}

func New(gps GPS, together Together) *Handler {
	return &Handler{
		gps:      gps,
		together: together,
	}
}

func (h *Handler) ConfigureHandlers(api *operations.SpyAPI) {
	api.PostBbinputHandler = operations.PostBbinputHandlerFunc(h.PostBbinputHandler)
	api.PostBbsHandler = operations.PostBbsHandlerFunc(h.PostBbsHandler)
}
