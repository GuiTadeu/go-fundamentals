package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	myStubSearchEngine := StubSearchEngine{}
	engine := NewStubEngine(myStubSearchEngine)
	expectedResult := []StubProduct{
		{"Calça para uma jovem", 350.00},
		{"Disco do O Terno", 66.00},
	}

	result := engine.StubGetAll()
	assert.Equal(t, expectedResult, result)
}