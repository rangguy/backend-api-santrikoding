package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"rangguy/backend-api/config"
	"strings"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil header authorization
		tokenString := c.Request.Header.Get("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token is required",
			})
			c.Abort() // menghentikan request selanjutnya
			return
		}

		// menghapus prefix bearer
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// struct untuk menampung klaim token
		claims := &jwt.RegisteredClaims{}

		// Parse token dan verifikasi tanda tangan dengan jwtKey
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Kembalikan kunci rahasia untuk memverifikasi token
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token is invalid",
			})
			c.Abort()
			return
		}

		// menyimpan klaim "sub" (username) ke dalam contract
		c.Set("username", claims.Subject)

		// lanjut ke handler selanjutnya
		c.Next()
	}
}
