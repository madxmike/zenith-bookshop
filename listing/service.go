package listing

import (
	"context"
	"fmt"

	"github.com/madxmike/zenith-bookshop/domain"
)

type Store interface {
	GetBookByISBN(context.Context, domain.ISBN) (domain.Book, error)
	AddListing(context.Context, Listing) error
	GetListingByISBN(context.Context, domain.ISBN) (Listing, error)
}

type Service struct {
	store Store
}

func NewService(store Store) Service {
	return Service{
		store: store,
	}
}

func (s *Service) ListBook(ctx context.Context, isbn domain.ISBN, price Price) error {
	if err := isbn.Validate(); err != nil {
		return fmt.Errorf("isbn is not valid: %w", err)
	}

	if err := price.Validate(); err != nil {
		return fmt.Errorf("price is not valid: %w", err)
	}

	if _, err := s.GetListing(ctx, isbn); err == nil {
		return fmt.Errorf("listing for %s already exists", isbn)
	}

	book, err := s.store.GetBookByISBN(ctx, isbn)
	if err != nil {
		return fmt.Errorf("could not get %s: %w", isbn, err)
	}

	listing := Listing{
		Book:  book,
		Price: price,
	}

	err = s.store.AddListing(ctx, listing)
	if err != nil {
		return fmt.Errorf("could not add listing for %s: %w", isbn, err)
	}

	return nil
}

func (s *Service) GetListing(ctx context.Context, ISBN domain.ISBN) (Listing, error) {
	listing, err := s.store.GetListingByISBN(ctx, ISBN)
	if err != nil {
		return Listing{}, fmt.Errorf("could not get listing for %s: %w", ISBN, err)
	}

	return listing, nil
}
