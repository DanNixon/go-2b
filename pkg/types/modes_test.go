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
