package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/madxmike/zenith-bookshop/http/rest"
	"github.com/madxmike/zenith-bookshop/library"
	"github.com/madxmike/zenith-bookshop/listing"
	"github.com/madxmike/zenith-bookshop/store/inmem"
)

func main() {

	store := inmem.NewStore()

	listingService := listing.NewService(&store)
	libraryService := library.NewService(&store)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/books", rest.NewLibraryHandler(libraryService))
	r.Mount("/listings", rest.NewListingHandler(listingService))

	http.ListenAndServe(":3000", r)
}
