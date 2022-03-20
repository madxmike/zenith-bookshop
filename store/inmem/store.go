package inmem

type Store struct {
	bookStore
	listingStore
}

func NewStore() Store {
	return Store{
		bookStore:    newBookStore(),
		listingStore: newListingStore(),
	}
}
