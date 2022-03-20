package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/madxmike/zenith-bookshop/domain"
	"github.com/madxmike/zenith-bookshop/library"
)

type libraryHandler struct {
	http.Handler
	service library.Service
}

func NewLibraryHandler(service library.Service) http.Handler {
	h := libraryHandler{
		service: service,
	}

	r := chi.NewMux()
	r.Post("/", h.PostBook)

	r.Route("/{isbn}", func(r chi.Router) {
		r.Get("/", h.GetBook)
		r.Delete("/", h.DeleteBook)
	})

	return r
}

func (h *libraryHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	isbn := domain.ISBN(chi.URLParam(r, "isbn"))

	book, err := h.service.GetBook(r.Context(), isbn)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(book); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *libraryHandler) PostBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.AddBook(r.Context(), book); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *libraryHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	isbn := domain.ISBN(chi.URLParam(r, "isbn"))

	if err := h.service.RemoveBook(r.Context(), isbn); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
