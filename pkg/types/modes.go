package types

import "errors"

type Mode string

const (
	ModePulse      = "pulse"
	ModeBounce     = "bounce"
	ModeContinuous = "continuous"
	ModeSplitA     = "split_a"
	ModeSplitB     = "split_b"
	ModeWave       = "wave"
	ModeWaterfall  = "waterfall"
	ModeSqueeze    = "squeeze"
	ModeMilk       = "milk"
	ModeThrob      = "throb"
	ModeThrust     = "thrust"
	ModeRandom     = "random"
	ModeStep       = "step"
	ModeTraining   = "training"
)

func AllModes() []Mode {
	return []Mode{ModePulse, ModeBounce, ModeContinuous, ModeSplitA, ModeSplitB, ModeWave, ModeWaterfall, ModeSqueeze, ModeMilk, ModeThrob, ModeThrust, ModeRandom, ModeStep, ModeTraining}
}

func (m Mode) Validate() error {
	for _, known := range AllModes() {
		if m == known {
			return nil
		}
	}
	return errors.New("Mode not valid")
}

var singleParameterModes = []Mode{ModeContinuous, ModeThrob, ModeThrust}

func (m Mode) ParameterCount() (int, error) {
	err := m.Validate()
	n := 2
	for _, sp := range singleParameterModes {
		if m == sp {
			n = 1
			break
		}
	}
	return n, err
}
