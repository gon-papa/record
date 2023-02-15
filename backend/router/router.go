package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewMux() http.Handler {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"res": "hello"}`))

	})
	return r
}
