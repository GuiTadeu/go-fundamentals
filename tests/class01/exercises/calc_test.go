package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	n1 := 3
	n2 := 5
	expectedResult := 8

	result := Sum(n1, n2)
	assert.Equal(t, expectedResult, result)
}

func TestSub(t *testing.T) {
	n1 := 5
	n2 := 3
	expectedResult := 2

	result := Sub(n1, n2)
	assert.Equal(t, expectedResult, result)
}

func TestDiv(t *testing.T) {
	n1 := 8
	n2 := 2
	expectedResult := 4

	result, err := Div(n1, n2)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestDivError(t *testing.T) {
	n1 := 8
	n2 := 0
	expectedResult := 0
	expectedError := "O denominador não pode ser 0"

	result, err := Div(n1, n2)

	assert.NotNil(t, err)
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedError, err.Error())
}