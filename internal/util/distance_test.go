package util

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/devchallenge/spy-api/internal/model"
)

func TestDistance(t *testing.T) {
	actual := Distance(model.Coordinate{139.74477, 35.6544}, model.Coordinate{139.8261, 34.4225})
	assert.InDelta(t, 137334.665, actual, 1)
}
