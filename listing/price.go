package listing

import "errors"

type Price int

func (cents Price) Validate() error {
	if cents < 0 {
		return errors.New("price cannot be less than 0")
	}

	return nil
}
