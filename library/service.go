package library

import (
	"context"
	"fmt"

	"github.com/madxmike/zenith-bookshop/domain"
)

type Store interface {
	AddBook(context.Context, domain.Book) error
	GetBookByISBN(context.Context, domain.ISBN) (domain.Book, error)
	RemoveBook(context.Context, domain.ISBN) error
}

type Service struct {
	store Store
}

func NewService(store Store) Service {
	return Service{
		store: store,
	}
}

func (s *Service) AddBook(ctx context.Context, b domain.Book) error {
	if err := b.Validate(); err != nil {
		return fmt.Errorf("book is not valid: %w", err)
	}

	if err := s.store.AddBook(ctx, b); err != nil {
		return fmt.Errorf("could not add book: %w", err)
	}

	return nil
}

func (s *Service) GetBook(ctx context.Context, isbn domain.ISBN) (domain.Book, error) {
	if err := isbn.Validate(); err != nil {
		return domain.Book{}, fmt.Errorf("isbn is not valid: %w", err)
	}

	book, err := s.store.GetBookByISBN(ctx, isbn)
	if err != nil {
		return domain.Book{}, fmt.Errorf("could not get book: %w", err)

	}

	return book, nil
}

func (s *Service) RemoveBook(ctx context.Context, isbn domain.ISBN) error {
	if err := isbn.Validate(); err != nil {
		return fmt.Errorf("isbn is not valid: %w", err)
	}

	if err := s.store.RemoveBook(ctx, isbn); err != nil {
		return fmt.Errorf("could not remove book: %w", err)
	}

	return nil
}
