package together

import (
	"testing"
	"time"

	"github.com/icrowley/fake"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/devchallenge/spy-api/internal/model"
)

//go:generate mockery -case=underscore -dir=. -outpkg=mock -output=./mock -recursive -all

func TestService_SpendPercentage(t *testing.T) {
	t.Run("invalid argument", func(t *testing.T) {
		s := &Service{}

		now := time.Now().UTC()
		actual, err := s.SpendPercentage(fake.Phone(), fake.Phone(), now.Add(time.Hour), now, 10)

		assert.Zero(t, actual)
		assert.Equal(t, model.ErrInvalidArgument, errors.Cause(err))

		actual, err = s.SpendPercentage(fake.Phone(), fake.Phone(), now, now.Add(time.Hour), 0)

		assert.Zero(t, actual)
		assert.Equal(t, model.ErrInvalidArgument, errors.Cause(err))
	})

	t.Run("basic case", func(t *testing.T) {
		t.SkipNow()
	})
}
