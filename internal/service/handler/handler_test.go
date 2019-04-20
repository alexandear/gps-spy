// Package handler_test contains integration tests of handler functions

package handler_test

import (
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
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

		bbinput(t, h, models.Number(fake.Phone()), fake.Longitude(), fake.Latitude())
		bbinput(t, h, models.Number(fake.Phone()), fake.Longitude(), fake.Latitude())
		bbinput(t, h, models.Number(fake.Phone()), fake.Longitude(), fake.Latitude())
	})
}

func TestHandler_PostBbsHandler(t *testing.T) {
	t.Run("when working hours", func(t *testing.T) {
		s := initStorage(t)
		defer util.Close(s)
		h := handler.New(gps.New(s), together.New(s))
		number1, number2 := models.Number(fake.Phone()), models.Number(fake.Phone())
		ts := models.Timestamp("2019/03/22-15:50:20")
		bbinputTs(t, h, number1, 22.1832284135991, 60.4538416572538, ts)
		bbinputTs(t, h, number2, 22.1832284135992, 60.4538416572539, ts)

		p := bbs(t, h, number1, number2,
			models.Timestamp("2019/01/01-00:00:00"), models.Timestamp("2019/12/31-00:00:00"), 10)

		assert.Zero(t, p)
	})

	t.Run("when not working hours", func(t *testing.T) {
		s := initStorage(t)
		defer util.Close(s)
		h := handler.New(gps.New(s), together.New(s))
		number1, number2 := models.Number(fake.Phone()), models.Number(fake.Phone())
		ts := models.Timestamp("2019/03/22-22:50:20")
		bbinputTs(t, h, number1, 22.1832284135991, 60.4538416572538, ts)
		bbinputTs(t, h, number2, 22.1832284135992, 60.4538416572539, ts)

		p := bbs(t, h, number1, number2,
			models.Timestamp("2019/01/01-00:00:00"), models.Timestamp("2019/12/31-00:00:00"), 100)

		assert.Equal(t, 100, p)
	})

	t.Run("when too far", func(t *testing.T) {
		s := initStorage(t)
		defer util.Close(s)
		h := handler.New(gps.New(s), together.New(s))
		number1, number2 := models.Number(fake.Phone()), models.Number(fake.Phone())
		ts := models.Timestamp("2019/03/22-22:50:20")
		bbinputTs(t, h, number1, 22.1832284135991, 60.4538416572538, ts)
		bbinputTs(t, h, number2, 40.1832284135992, 70.4538416572539, ts)

		p := bbs(t, h, number1, number2,
			models.Timestamp("2019/01/01-00:00:00"), models.Timestamp("2019/12/31-00:00:00"), 10)

		assert.Zero(t, p)
	})
}

func bbinput(t *testing.T, h *handler.Handler, number models.Number, lon, lat float32) {
	resp := h.PostBbinputHandler(operations.PostBbinputParams{Body: operations.PostBbinputBody{
		Imei:   stringPtr(fake.Characters()),
		Number: number,
		Coordinates: &operations.PostBbinputParamsBodyCoordinates{
			Longitude: &lon,
			Latitude:  &lat,
		},
	}})
	require.NotNil(t, resp)
	_, ok := resp.(*operations.PostBbinputOK)
	require.True(t, ok)
}

func bbinputTs(t *testing.T, h *handler.Handler, number models.Number, lon, lat float32, ts models.Timestamp) {
	resp := h.PostBbinputHandler(operations.PostBbinputParams{Body: operations.PostBbinputBody{
		Imei:   stringPtr(fake.Characters()),
		Number: number,
		Coordinates: &operations.PostBbinputParamsBodyCoordinates{
			Longitude: &lon,
			Latitude:  &lat,
		},
		Timestamp: ts,
	}})
	require.NotNil(t, resp)
	_, ok := resp.(*operations.PostBbinputOK)
	require.True(t, ok)
}

func bbs(t *testing.T, h *handler.Handler, number1, number2 models.Number, from, to models.Timestamp, minDistance int32) int {
	resp := h.PostBbsHandler(operations.PostBbsParams{
		Body: operations.PostBbsBody{
			Number1:     number1,
			Number2:     number2,
			From:        from,
			To:          to,
			MinDistance: minDistance,
		},
	})

	require.NotNil(t, resp)
	bbsOK, ok := resp.(*operations.PostBbsOK)
	require.True(t, ok)
	return int(bbsOK.Payload.Percentage)
}

func initStorage(t *testing.T) *storage.Storage {
	db, err := buntdb.Open(":memory:")
	require.NoError(t, err)
	return storage.New(db)
}

func stringPtr(val string) *string {
	return &val
}
