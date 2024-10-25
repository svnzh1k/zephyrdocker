package service

import (
	"errors"
	"time"
	"zephyr-api-mod/internal/models"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey string = "hello"
var tokenExpiration = time.Hour * 24

func GenerateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(tokenExpiration).Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(secretKey))
}

func ValidateJWT(token string) (*models.User, error) {
	claims := &jwt.MapClaims{}
	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, errors.New("invalid token")
	}

	user := &models.User{
		Id:       int((*claims)["id"].(float64)),
		Username: (*claims)["username"].(string),
		Role:     (*claims)["role"].(string),
	}

	return user, nil
}
