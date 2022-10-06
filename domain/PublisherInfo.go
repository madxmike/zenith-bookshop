package domain

import (
	"errors"
	"time"
)

type PublisherInfo struct {
	Published time.Time
	Publisher string
}

func (pi PublisherInfo) Validate() error {
	if pi.Published.IsZero() {
		return errors.New("published date must not be zero")
	}

	if time.Since(pi.Published) < 0 {
		return errors.New("published date must not be in the future")
	}

	if pi.Publisher == "" {
		return errors.New("publisher must not be empty")
	}

	return nil
}
