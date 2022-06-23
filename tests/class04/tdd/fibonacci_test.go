package tdd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci_5_numbers(t *testing.T) {
	expectedResult := []uint8{1, 1, 2, 3, 5}
	result := Fibonacci(5)
	assert.Equal(t, expectedResult, result)
}

func TestFibonacci_10_numbers(t *testing.T) {
	expectedResult := []uint8{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	result := Fibonacci(10)
	assert.Equal(t, expectedResult, result)
}