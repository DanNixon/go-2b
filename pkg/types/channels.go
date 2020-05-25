package types

type Channels struct {
	A      NumericSetting `json:"a"`
	B      NumericSetting `json:"b"`
	Linked bool           `json:"linked"`
}

func (c Channels) Validate() error {
	if err := c.A.Validate(); err != nil {
		return err
	}
	if err := c.B.Validate(); err != nil {
		return err
	}
	return nil
}
