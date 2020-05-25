package types

import "errors"

type PowerLevel string

const (
	PowerLevelLow  = "low"
	PowerLevelHigh = "high"
)

func AllPowerLevels() []PowerLevel {
	return []PowerLevel{PowerLevelLow, PowerLevelHigh}
}

func (pl PowerLevel) Validate() error {
	for _, known := range AllPowerLevels() {
		if pl == known {
			return nil
		}
	}
	return errors.New("Power level not valid")
}
