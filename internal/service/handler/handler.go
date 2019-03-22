package handler

import (
	"github.com/devchallenge/spy-api/internal/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) AddLocation(params operations.AddLocationParams) middleware.Responder {
	return nil
}

func (h *Handler) GetBboutputHandler(params operations.GetBboutputParams) middleware.Responder {
	return nil
}

func (h *Handler) ConfigureHandlers(api *operations.SpyAPI) {
	api.AddLocationHandler = operations.AddLocationHandlerFunc(h.AddLocation)
}
