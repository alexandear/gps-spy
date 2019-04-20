// Package handler_test contains helper functions for integration tests

package handler_test

import (
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/buntdb"

	"github.com/devchallenge/spy-api/internal/gen/models"
	"github.com/devchallenge/spy-api/internal/gen/restapi/operations"
	"github.com/devchallenge/spy-api/internal/service/handler"
	"github.com/devchallenge/spy-api/internal/storage"
)

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
			MinDistance: &minDistance,
		},
	})

	require.NotNil(t, resp)
	bbsOK, ok := resp.(*operations.PostBbsOK)
	require.True(t, ok)
	return int(*bbsOK.Payload.Percentage)
}

func initStorage(t *testing.T) *storage.Storage {
	db, err := buntdb.Open(":memory:")
	require.NoError(t, err)
	return storage.New(db)
}

func stringPtr(val string) *string {
	return &val
}
