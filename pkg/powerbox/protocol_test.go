package powerbox_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DanNixon/go-2b/pkg/powerbox"
	"github.com/DanNixon/go-2b/pkg/types"
)

func TestGenerateDeltaCommands1(t *testing.T) {
	s := types.Settings{
		PowerLevel: types.PowerLevelLow,
		Mode:       types.ModeBounce,
		Channels: types.Channels{
			A:      20,
			B:      40,
			Linked: false,
		},
		Parameters: types.Parameters{
			C: 45,
			D: 70,
		},
	}

	c, err := powerbox.GenerateDeltaCommands(types.Settings{}, s)
	assert.Nil(t, err)
	assert.Equal(t, 6, len(c))
	assert.Equal(t, "L", c[0])
	assert.Equal(t, "M1", c[1])
	assert.Equal(t, "C45", c[2])
	assert.Equal(t, "D70", c[3])
	assert.Equal(t, "A20", c[4])
	assert.Equal(t, "B40", c[5])
}

func TestGenerateDeltaCommands2(t *testing.T) {
	s := types.Settings{
		PowerLevel: types.PowerLevelHigh,
		Mode:       types.ModeTraining,
		Channels: types.Channels{
			A:      30,
			B:      0,
			Linked: true,
		},
		Parameters: types.Parameters{
			C: 20,
			D: 80,
		},
	}

	c, err := powerbox.GenerateDeltaCommands(types.Settings{}, s)
	assert.Nil(t, err)
	assert.Equal(t, 6, len(c))
	assert.Equal(t, "H", c[0])
	assert.Equal(t, "J1", c[1])
	assert.Equal(t, "M13", c[2])
	assert.Equal(t, "C20", c[3])
	assert.Equal(t, "D80", c[4])
	assert.Equal(t, "A30", c[5])
}

func TestParseStatusMessageValid(t *testing.T) {
	msg := "512:90:20:50:150:1:L:0:2.105"
	s, err := powerbox.ParseStatusMessage(msg)
	assert.Nil(t, err)
	expected := types.Status{
		Settings: types.Settings{
			PowerLevel: types.PowerLevelLow,
			Mode:       types.ModeBounce,
			Channels: types.Channels{
				A:      45,
				B:      10,
				Linked: false,
			},
			Parameters: types.Parameters{
				C: 25,
				D: 75,
			},
		},
		BatteryLevel:    512,
		FirmwareVersion: "2.105",
	}
	assert.Equal(t, expected, s)
}
