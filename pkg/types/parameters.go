package types

type Parameters struct {
	C NumericSetting `json:"c"`
	D NumericSetting `json:"d"`
}

func (p Parameters) Validate() error {
	if err := p.C.Validate(); err != nil {
		return err
	}
	if err := p.D.Validate(); err != nil {
		return err
	}
	return nil
}
