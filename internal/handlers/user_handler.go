package handlers

import (
	"api-service/internal/types"
	"api-service/pkg"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	
	switch r.Method {
	case http.MethodPost:
		var newUser types.User

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(data, &newUser)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		//validate
		if err := pkg.Validate(newUser); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
		if err := h.service.UserService.CreateUser(&newUser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
}
