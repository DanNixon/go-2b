package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DanNixon/go-2b/pkg/types"
)

func TestPowerLevelValidateValid(t *testing.T) {
	pl := types.PowerLevel("low")
	err := pl.Validate()
	assert.Nil(t, err)
}

func TestPowerLevelValidateInvalid(t *testing.T) {
	pl := types.PowerLevel("nope")
	err := pl.Validate()
	assert.NotNil(t, err)
}
