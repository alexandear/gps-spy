package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"

	"github.com/devchallenge/spy-api/internal/gen/restapi/operations"
	"github.com/devchallenge/spy-api/internal/model"
)

func (h *Handler) PostBbfastDriveHandler(params operations.PostBbfastDriveParams) middleware.Responder {
	switch numbers, err := h.violator.Numbers(); errors.Cause(err) {
	case nil:
		return operations.NewPostBbfastDriveOK().WithPayload(&operations.PostBbfastDriveOKBody{
			Phones: numbers,
		})
	case model.ErrInvalidArgument:
		return operations.NewPostBbfastDriveBadRequest()
	default:
		return operations.NewPostBbfastDriveInternalServerError()
	}
}
