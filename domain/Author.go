package domain

import "errors"

type Author struct {
	ID   int
	Name string
}

func (a Author) Validate() error {
	if a.Name == "" {
		return errors.New("author name must not be empty")
	}

	return nil
}
