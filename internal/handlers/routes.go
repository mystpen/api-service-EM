package handlers

import "net/http"

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	// mux.HandleFunc("/", h.home)

	return mux
}
