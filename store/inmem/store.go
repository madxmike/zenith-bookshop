package inmem

import (
	"context"
	"fmt"
	"sync"

	"github.com/madxmike/zenith-bookshop/domain"
	"github.com/madxmike/zenith-bookshop/listing"
)

type Store struct {
	sync.Mutex
	books    map[domain.ISBN]domain.Book
	listings map[domain.ISBN]listing.Listing
}

func NewStore() Store {
	return Store{
		books:    make(map[domain.ISBN]domain.Book),
		listings: make(map[domain.ISBN]listing.Listing),
	}
}

func (s *Store) GetBookByISBN(ctx context.Context, isbn domain.ISBN) (domain.Book, error) {
	s.Lock()
	defer s.Unlock()

	book, ok := s.books[isbn]
	if !ok {
		return domain.Book{}, fmt.Errorf("no book with isbn %s found", isbn)
	}

	return book, nil
}

func (s *Store) AddListing(ctx context.Context, listing listing.Listing) error {
	s.Lock()
	defer s.Unlock()

	isbn := listing.Book.ISBN
	if _, ok := s.listings[isbn]; !ok {
		return fmt.Errorf("listing of book with isbn %s already exists", isbn)
	}

	s.listings[isbn] = listing
	return nil
}

func (s *Store) GetListingByISBN(ctx context.Context, isbn domain.ISBN) (listing.Listing, error) {
	s.Lock()
	defer s.Unlock()

	_listing, ok := s.listings[isbn]
	if !ok {
		return listing.Listing{}, fmt.Errorf("no listing with isbn %s found", isbn)
	}

	return _listing, nil
}
