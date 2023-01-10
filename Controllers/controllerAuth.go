package controllers

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func createJWT(userID uint) (string, error) {
	// Buat token claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	// Buat token dengan metode HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Encode JWT menjadi string, sekaligus menyertakan secret key
	return token.SignedString([]byte("secret"))
}
