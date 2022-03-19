package domain

import "time"

type PublisherInfo struct {
	Published time.Time
	Publisher string
}

func (pi PublisherInfo) Validate() error {
	// TODO (Michael): We do not have any requirements for this right now
	return nil
}
