package handler

import (
	"fmt"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"

	"github.com/devchallenge/spy-api/internal/model"

	"github.com/devchallenge/spy-api/internal/gen/restapi/operations"
)

func (h *Handler) PostBbsHandler(params operations.PostBbsParams) middleware.Responder {
	body := params.Body
	if body.Number1 == "" {
		return newPostBbsBadRequest("number1 is required")
	}
	number1 := string(body.Number1)
	if body.Number2 == "" {
		return newPostBbsBadRequest("number2 is required")
	}
	number2 := string(body.Number2)
	from := time.Time{}
	if ts, err := ParseTimestamp(string(body.From)); err != nil {
		return newPostBbsBadRequest(fmt.Sprintf("failed to parse from, err=%v", err))
	} else {
		from = ts
	}
	to := time.Time{}
	if ts, err := ParseTimestamp(string(body.To)); err != nil {
		return newPostBbsBadRequest(fmt.Sprintf("failed to parse to, err=%v", err))
	} else {
		to = ts
	}
	distance := int(*body.MinDistance)
	switch p, err := h.together.SpendPercentage(number1, number2, from, to, distance); errors.Cause(err) {
	case nil:
		percentage := int32(p)
		return operations.NewPostBbsOK().WithPayload(&operations.PostBbsOKBody{
			Percentage: &percentage,
		})
	case model.ErrInvalidArgument:
		return newPostBbsBadRequest(err.Error())
	default:
		return operations.NewPostBbsInternalServerError().WithPayload(newError("failed to get spend percentage"))
	}
}

func newPostBbsBadRequest(message string) *operations.PostBbsBadRequest {
	return operations.NewPostBbsBadRequest().WithPayload(newError(message))
}
