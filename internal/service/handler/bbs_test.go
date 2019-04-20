package handler_test

import (
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"

	"github.com/devchallenge/spy-api/internal/gen/models"
	"github.com/devchallenge/spy-api/internal/service/gps"
	"github.com/devchallenge/spy-api/internal/service/handler"
	"github.com/devchallenge/spy-api/internal/service/together"
	"github.com/devchallenge/spy-api/internal/util"
)

func TestHandler_PostBbsHandler(t *testing.T) {
	t.Run("when working hours", func(t *testing.T) {
		s := initStorage(t)
		defer util.Close(s)
		h := handler.New(gps.New(s), together.New(s), nil)
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
		h := handler.New(gps.New(s), together.New(s), nil)
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
		h := handler.New(gps.New(s), together.New(s), nil)
		number1, number2 := models.Number(fake.Phone()), models.Number(fake.Phone())
		ts := models.Timestamp("2019/03/22-22:50:20")
		bbinputTs(t, h, number1, 22.1832284135991, 60.4538416572538, ts)
		bbinputTs(t, h, number2, 40.1832284135992, 70.4538416572539, ts)

		p := bbs(t, h, number1, number2,
			models.Timestamp("2019/01/01-00:00:00"), models.Timestamp("2019/12/31-00:00:00"), 10)

		assert.Zero(t, p)
	})
}
