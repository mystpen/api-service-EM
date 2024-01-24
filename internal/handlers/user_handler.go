package handlers

import (
	"api-service/internal/types"
	"api-service/pkg"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		
		filters := types.Filter{}

		page , err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil || page < 1 {
			page = 1
		}
		filters.Page = page

		age, err := strconv.Atoi(r.URL.Query().Get("age"))
		if err != nil || age < 1 {
			age = -1
		}
		filters.Age = age

		users, err := h.service.UserService.GetAllUsers(filters)

		jsonData, err := json.Marshal(users)
		if err != nil {
			http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var newUser types.User

		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(data, &newUser)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		// validate
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
	switch r.Method {
	case http.MethodDelete:
		user_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
		if err != nil || user_id <= 0{
			http.Error(w, "Invalid data", http.StatusBadRequest)
		}

		if err := h.service.UserService.DeleteUser(user_id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		user_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
		if err != nil || user_id <= 0{
			http.Error(w, "Invalid data", http.StatusBadRequest)
		}

		var updatedUser types.User

		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(data, &updatedUser)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		err = h.service.UserService.UpdateUser(&updatedUser, user_id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
