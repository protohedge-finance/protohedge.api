package http_transports

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewHealthHTTPHandler() http.Handler {
	m := mux.NewRouter()

	m.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	return m
}
