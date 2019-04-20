// Package handler_test contains integration tests of handler functions

package handler_test

import (
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/buntdb"

	"github.com/devchallenge/spy-api/internal/gen/models"
	"github.com/devchallenge/spy-api/internal/gen/restapi/operations"
	"github.com/devchallenge/spy-api/internal/service/gps"
	"github.com/devchallenge/spy-api/internal/service/handler"
	"github.com/devchallenge/spy-api/internal/service/together"
	"github.com/devchallenge/spy-api/internal/storage"
	"github.com/devchallenge/spy-api/internal/util"
)

func TestHandler_PostBbinputHandlerIntegration(t *testing.T) {
	t.Run("basic case", func(t *testing.T) {
		s := initStorage(t)
		defer util.Close(s)
		h := handler.New(gps.New(s), together.New(s))

		bbinput(t, h, fake.Phone(), fake.Longitude(), fake.Latitude())
		bbinput(t, h, fake.Phone(), fake.Longitude(), fake.Latitude())
		bbinput(t, h, fake.Phone(), fake.Longitude(), fake.Latitude())
	})
}

func TestHandler_PostBbsHandler(t *testing.T) {
	t.Run("basic case", func(t *testing.T) {
		s := initStorage(t)
		defer util.Close(s)
		h := handler.New(gps.New(s), together.New(s))

		number1 := fake.Phone()
		number2 := fake.Phone()

		bbinput(t, h, number1, 22.1832284135991, 60.4538416572538)
		bbinput(t, h, number2, 22.1832284135992, 60.4538416572539)
	})
}

func bbinput(t *testing.T, h *handler.Handler, number string, lon, lat float32) {
	resp := h.PostBbinputHandler(operations.PostBbinputParams{Body: operations.PostBbinputBody{
		Imei:   stringPtr(fake.Characters()),
		Number: models.Number(number),
		Coordinates: &operations.PostBbinputParamsBodyCoordinates{
			Longitude: &lon,
			Latitude:  &lat,
		},
	}})
	require.NotNil(t, resp)
	_, ok := resp.(*operations.PostBbinputOK)
	require.True(t, ok)
}

func initStorage(t *testing.T) *storage.Storage {
	db, err := buntdb.Open(":memory:")
	require.NoError(t, err)
	return storage.New(db)
}

func stringPtr(val string) *string {
	return &val
}
