package handler_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/devchallenge/spy-api/internal/gen/models"
	"github.com/devchallenge/spy-api/internal/gen/restapi/operations"
	"github.com/devchallenge/spy-api/internal/service/gps"
	"github.com/devchallenge/spy-api/internal/service/handler"
	"github.com/devchallenge/spy-api/internal/service/specnomery"
	"github.com/devchallenge/spy-api/internal/service/together"
	"github.com/devchallenge/spy-api/internal/service/violator"
	"github.com/devchallenge/spy-api/internal/util"
)

func TestHandler_PostBbfastDriveHandler(t *testing.T) {
	t.Run("invalid argument", func(t *testing.T) {
		s := initStorage(t)
		defer util.Close(s)
		h := handler.New(gps.New(s), together.New(s), violator.New(specnomery.EmptyAllowedUsersClient{}))

		resp := h.PostBbfastDriveHandler(operations.PostBbfastDriveParams{
			Body: operations.PostBbfastDriveBody{
				From:     models.Timestamp(""),
				To:       models.Timestamp("2019/03/22-15:50:20"),
				MinSpped: 110,
			},
		})

		require.NotNil(t, resp)
		_, ok := resp.(*operations.PostBbfastDriveBadRequest)
		require.True(t, ok)
	})
}
