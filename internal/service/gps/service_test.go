package gps

import (
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"

	"github.com/devchallenge/spy-api/internal/model"
	gpsMock "github.com/devchallenge/spy-api/internal/service/gps/mock"
)

//go:generate mockery -case=underscore -dir=. -outpkg=mock -output=./mock -recursive -all

func TestService_Add(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		sm := &gpsMock.Storage{}
		sm.On("Save", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		s := New(sm)

		err := s.Add(model.Phone{}, model.Coordinate{}, time.Time{})

		assert.NoError(t, err)
		sm.AssertExpectations(t)
	})

	t.Run("failure", func(t *testing.T) {
		sm := &gpsMock.Storage{}
		sm.On("Save", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("")).Once()
		s := New(sm)

		err := s.Add(model.Phone{}, model.Coordinate{}, time.Time{})

		assert.Error(t, err)
		sm.AssertExpectations(t)
	})
}
