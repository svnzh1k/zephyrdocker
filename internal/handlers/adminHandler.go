package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"zephyr-api-mod/internal/models"
	"zephyr-api-mod/internal/service"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)
	if user.Username == "" || user.Password == "" || len(user.Password) > 30 || len(user.Password) < 3 || len(user.Role) < 3 {
		http.Error(w, "a property for registration is missing", http.StatusBadRequest)
		return
	}
	if user.Role == "admin" && r.Context().Value("user").(*models.User).Role != "owner" {
		http.Error(w, "You need to have a role of Owner to add an admin", http.StatusUnauthorized)
		return
	}
	err := service.AddUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Successfully added a new user with role = " + user.Role))
}

func RemoveUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.RemoveUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Write([]byte("User successfully deleted"))
}

func GetWaitersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := service.GetWaiters()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	usersJson, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(usersJson)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := service.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	userJson, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(userJson)
}
