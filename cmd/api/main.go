package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/madxmike/zenith-bookshop/http/rest"
	"github.com/madxmike/zenith-bookshop/listing"
	"github.com/madxmike/zenith-bookshop/store/inmem"
)

func main() {

	store := inmem.NewStore()

	listingService := listing.NewService(&store)

	r := chi.NewRouter()
	r.Mount("/listing", rest.NewListingHandler(listingService))

	http.ListenAndServe(":3000", r)
}
