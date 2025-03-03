package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateDegrees(t *testing.T) {
	modifier := 15
	dc := 20

	result := CalculateDegrees(dc, modifier)

	assert.Equal(t, 0, result["critical_failures"])
	assert.Equal(t, 0, result["failures"])
	assert.Equal(t, 20, result["successes"])
	assert.Equal(t, 0, result["critical_successes"])
}
