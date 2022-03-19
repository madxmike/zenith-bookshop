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

func (s *Service) ListBook(ctx context.Context, ISBN domain.ISBN, price Price) error {
	if _, err := s.GetListing(ctx, ISBN); err == nil {
		return fmt.Errorf("listing for %s already exists", ISBN)
	}

	book, err := s.store.GetBookByISBN(ctx, ISBN)
	if err != nil {
		return fmt.Errorf("could not get %s: %w", ISBN, err)
	}

	listing := Listing{
		Book:  book,
		Price: price,
	}

	err = s.store.AddListing(ctx, listing)
	if err != nil {
		return fmt.Errorf("could not add listing for %s: %w", ISBN, err)
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
