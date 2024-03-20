package main

import (
	"net/http"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("location", r.URL.Path[1:])
	w.WriteHeader(http.StatusMovedPermanently)
}

func main() {
	handler := new(Handler)
	http.ListenAndServe(":8080", handler)
}
