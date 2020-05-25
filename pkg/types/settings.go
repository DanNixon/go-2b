package types

type Settings struct {
	PowerLevel PowerLevel `json:"power_level"`
	Mode       Mode       `json:"mode"`
	Channels   Channels   `json:"channels"`
	Parameters Parameters `json:"parameters"`
}

func (s Settings) Validate() error {
	if err := s.PowerLevel.Validate(); err != nil {
		return err
	}
	if err := s.Mode.Validate(); err != nil {
		return err
	}
	if err := s.Channels.Validate(); err != nil {
		return err
	}
	if err := s.Parameters.Validate(); err != nil {
		return err
	}
	return nil
}
