package handler

import (
	"net"
	"time"

	"github.com/devchallenge/spy-api/internal/restapi/operations"
	"github.com/devchallenge/spy-api/internal/service/model"
	"github.com/getfider/fider/app/pkg/errors"
	"github.com/go-openapi/runtime/middleware"
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

func (h *Handler) AddLocation(params operations.AddLocationParams) middleware.Responder {
	body := params.Body
	if body.Number == nil {
		return newAddLocationBadRequest("number is required")
	}
	if body.Imei == nil {
		return newAddLocationBadRequest("IMEI is required")
	}
	if body.Coordinates == nil {
		return newAddLocationBadRequest("coordinates are required")
	}
	if body.Coordinates.Longitude == nil {
		return newAddLocationBadRequest("longitude in coordinates is required")
	}
	if body.Coordinates.Latitude == nil {
		return newAddLocationBadRequest("latitude in coordinates is required")
	}
	location, err := model.NewLocation(*body.Number, *body.Imei, *body.Coordinates.Longitude, *body.Coordinates.Latitude)
	if err != nil {
		return newAddLocationBadRequest(err.Error())
	}
	if body.IP != "" {
		ip := net.ParseIP(body.IP)
		if ip == nil {
			return newAddLocationBadRequest("ip must be valid")
		}
		location.SetIP(ip)
	}
	if body.Timestamp != "" {
		timestamp, err := time.Parse("2006/01/02-15:04:05", body.Timestamp)
		if err != nil {
			return newAddLocationBadRequest("timestamp must be in format 'YYYY/MM/DD-hh:mm:ss'")
		}
		location.SetTimestamp(timestamp)
	}
	if err := h.storage.SaveLocation(location); err != nil {
		return newAddLocationServerError(errors.Wrap(err, "failed to save location"))
	}
	return operations.NewAddLocationOK()
}

func newAddLocationBadRequest(message string) *operations.AddLocationBadRequest {
	return operations.NewAddLocationBadRequest().WithPayload(model.NewError(message))
}

func newAddLocationServerError(err error) *operations.AddLocationInternalServerError {
	return operations.NewAddLocationInternalServerError().WithPayload(model.NewError(err.Error()))
}

func (h *Handler) GetBboutputHandler(params operations.GetBboutputParams) middleware.Responder {
	return middleware.NotImplemented("this endpoint is not implemented yet")
}

func (h *Handler) ConfigureHandlers(api *operations.SpyAPI) {
	api.AddLocationHandler = operations.AddLocationHandlerFunc(h.AddLocation)
}
