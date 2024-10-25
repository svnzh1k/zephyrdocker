package handlers

import (
	"encoding/json"
	"net/http"
	"zephyr-api-mod/internal/service"

	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid request body",
		})
		return
	}

	username, exists := data["username"]
	if !exists {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Missing username property",
		})
		return
	}

	password, exists := data["password"]
	if !exists {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Missing password property",
		})
		return
	}

	user, err := service.GetUserByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "No such user",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Password is incorrect",
		})
		return
	}

	token, err := service.GenerateJWT(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to generate token",
		})
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"username": user.Username,
		"role":     user.Role,
		"token":    "Bearer " + token,
	})
}
