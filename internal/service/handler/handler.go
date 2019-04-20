package handler

import (
	"net"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"

	"github.com/devchallenge/spy-api/internal/gen/models"
	"github.com/devchallenge/spy-api/internal/gen/restapi/operations"
	"github.com/devchallenge/spy-api/internal/model"
)

type Handler struct {
	gps GPS
}

type GPS interface {
	Add(phone model.Phone, coordinate model.Coordinate, timestamp time.Time) error
}

func New(gps GPS) *Handler {
	return &Handler{
		gps: gps,
	}
}

func (h *Handler) PostBbinputHandler(params operations.PostBbinputParams) middleware.Responder {
	body := params.Body
	if body.Number == nil {
		return newPostBbinputBadRequest("number is required")
	}
	if body.Imei == nil {
		return newPostBbinputBadRequest("IMEI is required")
	}
	ip := net.IP{}
	if body.IP != "" {
		if ip = net.ParseIP(body.IP); ip == nil {
			return newPostBbinputBadRequest("ip must be valid")
		}
	}
	phone, err := model.NewPhone(*body.Number, ip, *body.Imei)
	if err != nil {
		return newPostBbinputBadRequest(err.Error())
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
	coordinate, err := model.NewCoordinate(*body.Coordinates.Longitude, *body.Coordinates.Latitude)
	if err != nil {
		return newPostBbinputBadRequest(err.Error())
	}
	timestamp := time.Time{}
	if string(body.Timestamp) != "" {
		t, err := time.Parse("2006/01/02-15:04:05", string(body.Timestamp))
		if err != nil {
			return newPostBbinputBadRequest("timestamp must be in format 'YYYY/MM/DD-hh:mm:ss'")
		}
		loc, err := time.LoadLocation("Europe/Kiev")
		if err != nil {
			return newPostBbinputServerError(errors.Wrap(err, "failed to load location"))
		}
		timestamp = t.In(loc)
	}
	switch err := h.gps.Add(phone, coordinate, timestamp); errors.Cause(err) {
	case nil:
		return operations.NewPostBbinputOK()
	case model.ErrInvalidArgument:
		return newPostBbinputBadRequest(err.Error())
	default:
		return newPostBbinputServerError(err)
	}
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
