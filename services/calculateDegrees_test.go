package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateDegrees(t *testing.T) {
	t.Run("sucess - Provided example", func(t *testing.T) {
		modifier := 15
		dc := 20

		result := CalculateDegrees(dc, modifier)

		assert.Equal(t, 1, result["critical_failures"])
		assert.Equal(t, 3, result["failures"])
		assert.Equal(t, 10, result["successes"])
		assert.Equal(t, 6, result["critical_successes"])
	})

	t.Run("sucess - Modifier > DC", func(t *testing.T) {
		modifier := 20
		dc := 10

		result := CalculateDegrees(dc, modifier)

		assert.Equal(t, 0, result["critical_failures"])
		assert.Equal(t, 0, result["failures"])
		assert.Equal(t, 1, result["successes"])
		assert.Equal(t, 19, result["critical_successes"])
	})

	t.Run("sucess - Modifier = DC", func(t *testing.T) {
		modifier := 10
		dc := 10

		result := CalculateDegrees(dc, modifier)

		assert.Equal(t, 0, result["critical_failures"])
		assert.Equal(t, 1, result["failures"])
		assert.Equal(t, 8, result["successes"])
		assert.Equal(t, 11, result["critical_successes"])
	})
}
