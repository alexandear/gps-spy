package handler

import (
	"net"
	"time"

	"github.com/getfider/fider/app/pkg/errors"
	"github.com/go-openapi/runtime/middleware"

	"github.com/devchallenge/spy-api/internal/gen/models"
	"github.com/devchallenge/spy-api/internal/gen/restapi/operations"
	"github.com/devchallenge/spy-api/internal/model"
)

type Handler struct {
	storage Storage
}

type Storage interface {
	SaveLocation(location *model.Location) error
}

func New(storage Storage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) PostBbinputHandler(params operations.PostBbinputParams) middleware.Responder {
	body := params.Body
	if body.Number == nil {
		return newPostBbinputBadRequest("number is required")
	}
	if body.Imei == nil {
		return newPostBbinputBadRequest("IMEI is required")
	}
	if body.Coordinates == nil {
		return newPostBbinputBadRequest("coordinates are required")
	}
	if body.Coordinates.Longitude == nil {
		return newPostBbinputBadRequest("longitude in coordinates is required")
	}
	if body.Coordinates.Latitude == nil {
		return newPostBbinputBadRequest("latitude in coordinates is required")
	}
	location, err := model.NewLocation(*body.Number, *body.Imei, *body.Coordinates.Longitude, *body.Coordinates.Latitude)
	if err != nil {
		return newPostBbinputBadRequest(err.Error())
	}
	if body.IP != "" {
		ip := net.ParseIP(body.IP)
		if ip == nil {
			return newPostBbinputBadRequest("ip must be valid")
		}
		location.SetIP(ip)
	}
	if body.Timestamp != "" {
		timestamp, err := time.Parse("2006/01/02-15:04:05", body.Timestamp)
		if err != nil {
			return newPostBbinputBadRequest("timestamp must be in format 'YYYY/MM/DD-hh:mm:ss'")
		}
		location.SetTimestamp(timestamp)
	}
	if err := h.storage.SaveLocation(location); err != nil {
		return newPostBbinputServerError(errors.Wrap(err, "failed to save location"))
	}
	return operations.NewPostBbinputOK()
}

func newPostBbinputBadRequest(message string) *operations.PostBbinputBadRequest {
	return operations.NewPostBbinputBadRequest().WithPayload(newError(message))
}

func newPostBbinputServerError(err error) *operations.PostBbinputInternalServerError {
	return operations.NewPostBbinputInternalServerError().WithPayload(newError(err.Error()))
}

func (h *Handler) ConfigureHandlers(api *operations.SpyAPI) {
	api.PostBbinputHandler = operations.PostBbinputHandlerFunc(h.PostBbinputHandler)
}

func newError(message string) *models.Error {
	return &models.Error{
		Message: &message,
	}
}
