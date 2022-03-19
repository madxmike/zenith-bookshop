package domain

import (
	"errors"
)

type ISBN string

func (isbn ISBN) Validate() error {
	// Note (Michael): For our purposes right now, we are taking a simplified view of ISBN
	if len(isbn) != 10 && len(isbn) != 13 {
		return errors.New("isbn must be a length of 10 or 13")
	}

	return nil
}
