package listing

import "github.com/madxmike/zenith-bookshop/domain"

type Listing struct {
	ID    int
	Book  domain.Book
	Price Price
}
