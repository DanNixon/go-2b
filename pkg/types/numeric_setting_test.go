package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DanNixon/go-2b/pkg/types"
)

func TestNumericSettingValidateInvalidUpper(t *testing.T) {
	ns := types.NumericSetting(101)
	err := ns.Validate()
	assert.NotNil(t, err)
}

func TestNumericSettingValidateValidUpper(t *testing.T) {
	ns := types.NumericSetting(100)
	err := ns.Validate()
	assert.Nil(t, err)
}
func TestNumericSettingValidateValidLower(t *testing.T) {
	ns := types.NumericSetting(0)
	err := ns.Validate()
	assert.Nil(t, err)
}

func TestNumericSettingValidateInvalidLower(t *testing.T) {
	ns := types.NumericSetting(-1)
	err := ns.Validate()
	assert.NotNil(t, err)
}
