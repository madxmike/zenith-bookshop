package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/madxmike/zenith-bookshop/domain"
	"github.com/madxmike/zenith-bookshop/listing"

	"github.com/go-chi/chi/v5"
)

type listingHandler struct {
	http.Handler
	service listing.Service
}

func NewListingHandler(service listing.Service) http.Handler {
	h := listingHandler{
		service: service,
	}

	r := chi.NewMux()
	r.Get("/{isbn}", h.GetListing)
	r.Post("/", h.ListBook)

	return r
}

func (h *listingHandler) ListBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var body struct {
		ISBN  domain.ISBN   `json:"ISBN"`
		Price listing.Price `json:"price"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := body.ISBN.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := body.Price.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.ListBook(r.Context(), body.ISBN, body.Price); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *listingHandler) GetListing(w http.ResponseWriter, r *http.Request) {
	isbn := domain.ISBN(chi.URLParam(r, "isbn"))

	if err := isbn.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	listing, err := h.service.GetListing(r.Context(), domain.ISBN(isbn))
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(listing); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
