package domain

import (
	"errors"
	"fmt"
)

type Book struct {
	ISBN          ISBN
	Title         string
	Authors       []Author
	Description   string
	Pages         int
	PublisherInfo PublisherInfo
	Language      string
}

func (b Book) Validate() error {
	if err := b.ISBN.Validate(); err != nil {
		return fmt.Errorf("book isbn must be valid: %w", err)
	}

	if b.Title == "" {
		return errors.New("book title must not be nil")
	}

	if len(b.Authors) == 0 {
		return errors.New("book must have atleast one author")
	}

	for _, author := range b.Authors {
		if err := author.Validate(); err != nil {
			return fmt.Errorf("book author must be valid: %w", err)
		}
	}

	if b.Pages <= 0 {
		return errors.New("book pages must be greater than 0")
	}

	if err := b.PublisherInfo.Validate(); err != nil {
		return fmt.Errorf("book publisher info must be valid: %w", err)
	}

	return nil
}
