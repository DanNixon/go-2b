package types

import (
	"errors"
	"fmt"
)

type NumericSetting int

func (n NumericSetting) Validate() error {
	if n < 0 || n > 100 {
		return errors.New(fmt.Sprintf("Numeric value (%d) out of range (0-100)", n))
	}
	return nil
}
