package inmem

import (
	"context"
	"fmt"
	"sync"

	"github.com/madxmike/zenith-bookshop/domain"
)

type bookStore struct {
	sync.Mutex
	books map[domain.ISBN]domain.Book
}

func newBookStore() bookStore {
	return bookStore{
		books: make(map[domain.ISBN]domain.Book),
	}
}

func (s *bookStore) GetBookByISBN(ctx context.Context, isbn domain.ISBN) (domain.Book, error) {
	s.Lock()
	defer s.Unlock()

	book, ok := s.books[isbn]
	if !ok {
		return domain.Book{}, fmt.Errorf("no book with isbn %s found", isbn)
	}

	return book, nil
}

func (s *bookStore) AddBook(ctx context.Context, b domain.Book) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.books[b.ISBN]; ok {
		return fmt.Errorf("book with isbn %s already exists", b.ISBN)
	}

	s.books[b.ISBN] = b

	return nil
}

func (s *bookStore) RemoveBook(ctx context.Context, isbn domain.ISBN) error {
	if _, ok := s.books[isbn]; !ok {
		return fmt.Errorf("no book with isbn %s found", isbn)
	}

	delete(s.books, isbn)
	return nil
}
