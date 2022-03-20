package inmem

import (
	"context"
	"fmt"
	"sync"

	"github.com/madxmike/zenith-bookshop/domain"
	"github.com/madxmike/zenith-bookshop/listing"
)

type listingStore struct {
	sync.Mutex

	listings map[domain.ISBN]listing.Listing
}

func newListingStore() listingStore {
	return listingStore{
		listings: make(map[domain.ISBN]listing.Listing),
	}
}

func (s *listingStore) AddListing(ctx context.Context, listing listing.Listing) error {
	s.Lock()
	defer s.Unlock()

	isbn := listing.Book.ISBN
	if _, ok := s.listings[isbn]; !ok {
		return fmt.Errorf("listing of book with isbn %s already exists", isbn)
	}

	s.listings[isbn] = listing
	return nil
}

func (s *listingStore) GetListingByISBN(ctx context.Context, isbn domain.ISBN) (listing.Listing, error) {
	s.Lock()
	defer s.Unlock()

	_listing, ok := s.listings[isbn]
	if !ok {
		return listing.Listing{}, fmt.Errorf("no listing with isbn %s found", isbn)
	}

	return _listing, nil
}
