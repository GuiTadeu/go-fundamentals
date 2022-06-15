package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateName(t *testing.T) {

	product := products[0]
	expectedOldName := "Before Update"

	assert.Equal(t, expectedOldName, product.name)

	myMockSearchEngine := MockSearchEngine{}
	engine := NewMockEngine(myMockSearchEngine)
	newName := "After Update"
	
	engine.UpdateName(1, newName)

	assert.Equal(t, newName, products[0].name)
	assert.True(t, engine.GetAllWasCalled)
}