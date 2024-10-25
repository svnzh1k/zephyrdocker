package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
	"zephyr-api-mod/internal/models"
	"zephyr-api-mod/internal/service"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func RequestLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var color string
		switch r.Method {
		case http.MethodGet:
			color = Blue
		case http.MethodPost:
			color = Green
		case http.MethodPut:
			color = Yellow
		case http.MethodDelete:
			color = Red
		default:
			color = Reset
		}

		fmt.Printf("[%s] %sMethod: %s%s | URL: %s\n",
			time.Now().Format("2006-01-02 15:04:05"),
			color,
			r.Method,
			Reset,
			r.URL.Path,
		)
		next(w, r)
	}
}

func TokenValidator(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		if len(authHeader) < 7 || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}

		token := authHeader[7:]

		user, err := service.ValidateJWT(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func AdminRoleValidator(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, exists := r.Context().Value("user").(*models.User)
		if !exists {
			http.Error(w, "No user in context", http.StatusInternalServerError)
			return
		}
		if user.Role != "admin" {
			http.Error(w, "You are not an admin", http.StatusInternalServerError)
			return
		}
		next(w, r)
	}
}

func OwnerRoleValidator(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, exists := r.Context().Value("user").(*models.User)
		if !exists {
			http.Error(w, "No user in context", http.StatusInternalServerError)
			return
		}
		if user.Role != "owner" {
			http.Error(w, "You are not an owner", http.StatusInternalServerError)
			return
		}
		next(w, r)
	}
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow all origins
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
