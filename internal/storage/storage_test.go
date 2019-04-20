package storage

import (
	"net"
	"testing"
	"time"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/buntdb"

	"github.com/devchallenge/spy-api/internal/model"
)

func TestStorage_Save(t *testing.T) {
	s := initStorage(t)

	t.Run("success when one number", func(t *testing.T) {
		number := fake.Phone()
		phone := model.Phone{
			Number: number,
			IMEI:   fake.CharactersN(10),
			IP:     net.ParseIP(fake.IPv4()),
		}
		coordinate := model.Coordinate{
			Longitude: fake.Longitude(),
			Latitude:  fake.Latitude(),
		}
		timestamp := time.Now().UTC()

		err := s.Save(phone, coordinate, timestamp)

		require.NoError(t, err)
		items, err := s.read(number)
		require.NoError(t, err)
		require.Len(t, items, 1)
		assert.Equal(t, &item{
			IP:         phone.IP,
			IMEI:       phone.IMEI,
			Coordinate: coordinate,
			Timestamp:  timestamp,
		}, items[0])
	})

	t.Run("success when few numbers", func(t *testing.T) {
		number := fake.Phone()
		phone := model.Phone{
			Number: number,
			IMEI:   fake.CharactersN(10),
			IP:     net.ParseIP(fake.IPv4()),
		}
		coordinate := model.Coordinate{
			Longitude: fake.Longitude(),
			Latitude:  fake.Latitude(),
		}
		timestamp := time.Now().UTC()
		require.NoError(t, s.Save(phone, coordinate, timestamp))
		require.NoError(t, s.Save(model.Phone{
			Number: fake.Phone(),
			IMEI:   fake.CharactersN(10),
			IP:     net.ParseIP(fake.IPv4()),
		}, model.Coordinate{
			Longitude: fake.Longitude(),
			Latitude:  fake.Latitude(),
		}, time.Now().UTC().Add(time.Hour)))

		items, err := s.read(number)

		require.NoError(t, err)
		require.Len(t, items, 1)
		assert.Equal(t, &item{
			IP:         phone.IP,
			IMEI:       phone.IMEI,
			Coordinate: coordinate,
			Timestamp:  timestamp,
		}, items[0])
	})
}

func TestStorage_read(t *testing.T) {
	s := initStorage(t)

	t.Run("success", func(t *testing.T) {
		number := fake.Phone()
		phone := model.Phone{
			Number: number,
			IMEI:   fake.CharactersN(10),
			IP:     net.ParseIP(fake.IPv4()),
		}
		coordinate := model.Coordinate{
			Longitude: fake.Longitude(),
			Latitude:  fake.Latitude(),
		}
		timestamp := time.Now().UTC()
		require.NoError(t, s.Save(phone, coordinate, timestamp))

		items, err := s.read(number)

		require.NoError(t, err)
		require.Len(t, items, 1)
		assert.Equal(t, &item{
			IP:         phone.IP,
			IMEI:       phone.IMEI,
			Coordinate: coordinate,
			Timestamp:  timestamp,
		}, items[0])
	})
}

func initStorage(t *testing.T) *Storage {
	db, err := buntdb.Open(":memory:")
	require.NoError(t, err)
	return New(db)
}
