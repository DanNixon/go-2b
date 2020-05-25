package powerbox

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/DanNixon/go-2b/pkg/types"
)

const (
	KillCommand  = "K"
	ResetCommand = "E"
)

var powerLevelToSerial = map[types.PowerLevel]string{
	types.PowerLevelLow:  "L",
	types.PowerLevelHigh: "H",
}

var powerLevelFromSerial = map[string]types.PowerLevel{
	"L": types.PowerLevelLow,
	"H": types.PowerLevelHigh,
}

var channelLinkedToSerial = map[bool]string{
	false: "0",
	true:  "1",
}

var channelLinkedFromSerial = map[string]bool{
	"0": false,
	"1": true,
}

var modeToSerial = map[types.Mode]string{
	types.ModePulse:      "0",
	types.ModeBounce:     "1",
	types.ModeContinuous: "2",
	types.ModeSplitA:     "3",
	types.ModeSplitB:     "4",
	types.ModeWave:       "5",
	types.ModeWaterfall:  "6",
	types.ModeSqueeze:    "7",
	types.ModeMilk:       "8",
	types.ModeThrob:      "9",
	types.ModeThrust:     "10",
	types.ModeRandom:     "11",
	types.ModeStep:       "12",
	types.ModeTraining:   "13",
}

var modeFromSerial = map[string]types.Mode{
	"0":  types.ModePulse,
	"1":  types.ModeBounce,
	"2":  types.ModeContinuous,
	"3":  types.ModeSplitA,
	"4":  types.ModeSplitB,
	"5":  types.ModeWave,
	"6":  types.ModeWaterfall,
	"7":  types.ModeSqueeze,
	"8":  types.ModeMilk,
	"9":  types.ModeThrob,
	"10": types.ModeThrust,
	"11": types.ModeRandom,
	"12": types.ModeStep,
	"13": types.ModeTraining,
}

func GenerateDeltaCommands(current types.Settings, target types.Settings) ([]string, error) {
	log.Printf("Current settings: %v", current)
	log.Printf("Target settings: %v", target)

	commands := []string{}

	if err := target.Validate(); err != nil {
		return commands, err
	}

	if target.PowerLevel != current.PowerLevel {
		commands = append(commands, fmt.Sprintf("%s", powerLevelToSerial[target.PowerLevel]))
		current = types.Settings{}
	}

	if target.Channels.Linked != current.Channels.Linked {
		commands = append(commands, fmt.Sprintf("J%s", channelLinkedToSerial[target.Channels.Linked]))
		current = types.Settings{}
	}

	if target.Mode != current.Mode {
		commands = append(commands, fmt.Sprintf("M%s", modeToSerial[target.Mode]))
		current = types.Settings{}
	}

	if target.Parameters.C != current.Parameters.C {
		commands = append(commands, fmt.Sprintf("C%d", target.Parameters.C))
	}

	if target.Parameters.D != current.Parameters.D {
		commands = append(commands, fmt.Sprintf("D%d", target.Parameters.D))
	}

	if target.Channels.A != current.Channels.A {
		commands = append(commands, fmt.Sprintf("A%d", target.Channels.A))
	}

	if target.Channels.B != current.Channels.B {
		commands = append(commands, fmt.Sprintf("B%d", target.Channels.B))
	}

	return commands, nil
}

func parseNumber(s string) (int, error) {
	v, err := strconv.ParseInt(s, 10, 32)
	return int(v), err
}

func ParseStatusMessage(msg string) (types.Status, error) {
	var s types.Status

	parts := strings.Split(msg, ":")
	if len(parts) != 9 {
		return s, errors.New(fmt.Sprintf("Incorrect number of parts (%d) in status string", len(parts)))
	}

	if v, err := parseNumber(parts[0]); err != nil {
		return s, err
	} else {
		s.BatteryLevel = v
	}

	if v, err := parseNumber(parts[1]); err != nil {
		return s, err
	} else {
		s.Settings.Channels.A = types.NumericSetting(v / 2)
	}

	if v, err := parseNumber(parts[2]); err != nil {
		return s, err
	} else {
		s.Settings.Channels.B = types.NumericSetting(v / 2)
	}

	if v, err := parseNumber(parts[3]); err != nil {
		return s, err
	} else {
		s.Settings.Parameters.C = types.NumericSetting(v / 2)
	}

	if v, err := parseNumber(parts[4]); err != nil {
		return s, err
	} else {
		s.Settings.Parameters.D = types.NumericSetting(v / 2)
	}

	s.Settings.Mode = modeFromSerial[parts[5]]
	s.Settings.PowerLevel = powerLevelFromSerial[parts[6]]
	s.Settings.Channels.Linked = channelLinkedFromSerial[parts[7]]

	s.FirmwareVersion = parts[8]

	return s, nil
}
