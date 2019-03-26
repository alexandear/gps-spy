package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/require"

	"github.com/devchallenge/spy-api/internal/gen/restapi/operations"
	"github.com/devchallenge/spy-api/internal/model"
)

func TestHandler_AddLocation(t *testing.T) {
	number := fake.Phone()
	imei := fake.CharactersN(10)
	longitude := fake.Longitude()
	latitude := fake.Latitude()
	wrongLongitude := float32(181.0)
	wrongLatitude := float32(91.0)

	t.Run("when invalid arguments returns 400", func(t *testing.T) {
		for name, tc := range map[string]struct {
			body operations.PostBbinputBody
		}{
			"empty number": {
				body: operations.PostBbinputBody{
					Number: nil,
					Imei:   &imei,
					Coordinates: &operations.PostBbinputParamsBodyCoordinates{
						Longitude: &longitude,
						Latitude:  &latitude,
					},
				},
			},
			"empty imei": {
				body: operations.PostBbinputBody{
					Number: &number,
					Imei:   nil,
					Coordinates: &operations.PostBbinputParamsBodyCoordinates{
						Longitude: &longitude,
						Latitude:  &latitude,
					},
				},
			},
			"empty coordinates": {
				body: operations.PostBbinputBody{
					Number:      &number,
					Imei:        &imei,
					Coordinates: nil,
				},
			},
			"empty longitude": {
				body: operations.PostBbinputBody{
					Number: &number,
					Imei:   &imei,
					Coordinates: &operations.PostBbinputParamsBodyCoordinates{
						Longitude: nil,
						Latitude:  &latitude,
					},
				},
			},
			"empty latitude": {
				body: operations.PostBbinputBody{
					Number: &number,
					Imei:   &imei,
					Coordinates: &operations.PostBbinputParamsBodyCoordinates{
						Longitude: &longitude,
						Latitude:  nil,
					},
				},
			},
			"wrong longitude": {
				body: operations.PostBbinputBody{
					Number: &number,
					Imei:   &imei,
					Coordinates: &operations.PostBbinputParamsBodyCoordinates{
						Longitude: &wrongLongitude,
						Latitude:  &latitude,
					},
				},
			},
			"wrong latitude": {
				body: operations.PostBbinputBody{
					Number: &number,
					Imei:   &imei,
					Coordinates: &operations.PostBbinputParamsBodyCoordinates{
						Longitude: &longitude,
						Latitude:  &wrongLatitude,
					},
				},
			},
			"wrong ip": {
				body: operations.PostBbinputBody{
					Number: &number,
					Imei:   &imei,
					Coordinates: &operations.PostBbinputParamsBodyCoordinates{
						Longitude: &longitude,
						Latitude:  &latitude,
					},
					IP: "300.300.300.300",
				},
			},
			"wrong timestamp": {
				body: operations.PostBbinputBody{
					Number: &number,
					Imei:   &imei,
					Coordinates: &operations.PostBbinputParamsBodyCoordinates{
						Longitude: &longitude,
						Latitude:  &latitude,
					},
					Timestamp: "02 Jan 06 15:04 MST",
				},
			},
		} {
			t.Run(name, func(t *testing.T) {
				h := New(&storageStub{})
				httpReq := http.Request{}
				params := operations.PostBbinputParams{
					HTTPRequest: httpReq.WithContext(context.Background()),
					Body:        tc.body,
				}

				resp := h.PostBbinputHandler(params)

				require.NotNil(t, resp)
				_, ok := resp.(*operations.PostBbinputBadRequest)
				require.True(t, ok)
			})
		}
	})
}

type storageStub struct {
}

func (s *storageStub) SaveLocation(*model.Location) error {
	return nil
}
