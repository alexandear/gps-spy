package handler

import (
	"fmt"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"

	"github.com/devchallenge/spy-api/internal/gen/restapi/operations"
	"github.com/devchallenge/spy-api/internal/model"
)

func (h *Handler) PostBbfastDriveHandler(params operations.PostBbfastDriveParams) middleware.Responder {
	body := params.Body
	from := time.Time{}
	if ts, err := ParseTimestamp(string(body.From)); err != nil {
		return newPostBbfastDriveBadRequest(fmt.Sprintf("failed to parse from, err=%v", err))
	} else {
		from = ts
	}
	to := time.Time{}
	if ts, err := ParseTimestamp(string(body.To)); err != nil {
		return newPostBbfastDriveBadRequest(fmt.Sprintf("failed to parse to, err=%v", err))
	} else {
		to = ts
	}
	minSpeed := int(body.MinSpped)
	minCoordinate, err := model.NewCoordinate(*body.MinLocation.Longitude, *body.MinLocation.Latitude)
	if err != nil {
		return newPostBbsBadRequest("failed to parse min location")
	}
	maxCoordinate, err := model.NewCoordinate(*body.MaxLocation.Longitude, *body.MaxLocation.Latitude)
	if err != nil {
		return newPostBbsBadRequest("failed to parse max location")
	}
	switch numbers, err := h.violator.Numbers(from, to, minSpeed,
		minCoordinate, maxCoordinate); errors.Cause(err) {
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

func newPostBbfastDriveBadRequest(message string) *operations.PostBbfastDriveBadRequest {
	return operations.NewPostBbfastDriveBadRequest().WithPayload(newError(message))
}
