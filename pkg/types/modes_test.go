package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DanNixon/go-2b/pkg/types"
)

func TestModeValidateValid(t *testing.T) {
	mode := types.Mode("pulse")
	err := mode.Validate()
	assert.Nil(t, err)
}

func TestModeValidateInvalid(t *testing.T) {
	mode := types.Mode("nope")
	err := mode.Validate()
	assert.NotNil(t, err)
}

func TestModeParameterCount1(t *testing.T) {
	mode := types.Mode("continuous")
	paramCount, err := mode.ParameterCount()
	assert.Nil(t, err)
	assert.Equal(t, paramCount, 1)
}

func TestModeParameterCount2(t *testing.T) {
	mode := types.Mode("pulse")
	paramCount, err := mode.ParameterCount()
	assert.Nil(t, err)
	assert.Equal(t, paramCount, 2)
}
