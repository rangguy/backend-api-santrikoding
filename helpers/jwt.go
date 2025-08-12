package helpers

import (
	"github.com/golang-jwt/jwt/v5"
	"rangguy/backend-api/config"
	"time"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

func GenerateToken(username string) string {
	// membuat expire time 60 menit
	expirationTime := time.Now().Add(60 * time.Minute)

	// membuat claim yang memuat subject username dan expiredat
	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	// membuat token baru dengan klaim yang telah dibuat
	// mengguakan algoritma HS256 untuk signed token
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)

	// Mengembalikan token dalam bentuk string
	return token
}
