package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateOK(t *testing.T) {
	entity := Task{
		ID:   1,
		Name: "Fernando",
	}

	err := entity.Validate()
	assert.Nil(t, err)
}

func TestValidateNameMin(t *testing.T) {
	entity := Task{
		ID:   1,
		Name: "Fern",
	}

	err := entity.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "name: the length must be between 5 and 50.")
}

func TestValidateNameMax(t *testing.T) {
	entity := Task{
		ID:   1,
		Name: "Fernando Fernando Fernando Fernando Fernando Fernando",
	}

	err := entity.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "name: the length must be between 5 and 50.")
}

func TestValidateNameRequired(t *testing.T) {
	entity := Task{
		ID: 1,
	}

	err := entity.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "name: cannot be blank.")
}
