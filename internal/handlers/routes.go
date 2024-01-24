package handlers

import "net/http"

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.GetAll)
	mux.HandleFunc("/create", h.Create)
	mux.HandleFunc("/delete", h.Delete)
	mux.HandleFunc("/update", h.Update)

	return mux
}
