package middleware

import (
	"log"
	"net/http"
	"strings"
	"ticketing-konser/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	JWTUtil *utils.JWTUtil
}

// NewAuthMiddleware membuat instance baru dari AuthMiddleware
func NewAuthMiddleware(jwtUtil *utils.JWTUtil) *AuthMiddleware {
	return &AuthMiddleware{JWTUtil: jwtUtil}
}

// Middleware memvalidasi token JWT dari header Authorization
func (a *AuthMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header tidak ditemukan"})
			c.Abort()
			return
		}

		// Validasi format Authorization header
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Format Authorization header salah"})
			c.Abort()
			return
		}

		// Ambil token dari header
		tokenString := parts[1]
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak boleh kosong"})
			c.Abort()
			return
		}

		// Validasi token menggunakan JWTUtil
		claims, err := a.JWTUtil.ValidateToken(tokenString)
		if err != nil {
			log.Printf("Token tidak valid: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("userRole", claims.Role)
		c.Next()
	}
}
