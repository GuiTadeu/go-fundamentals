package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsc(t *testing.T) {

	expectedResult := []int{2, 3, 4, 5, 6, 7, 8}
	result := Asc(5, 8, 3, 4, 7, 2, 6)

	assert.Equal(t, expectedResult, result)
}