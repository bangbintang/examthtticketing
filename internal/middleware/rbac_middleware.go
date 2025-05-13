package middleware

import (
	"net/http"
	"ticketing-konser/internal/utils"

	"github.com/gin-gonic/gin"
)

// Role constants
const (
	RoleAdmin    = "admin"
	RoleCustomer = "customer"
)

// RBACMiddleware membatasi akses berdasarkan peran yang diizinkan
func RBACMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := utils.DefaultLoggerInstance

		// Ambil role user dari context
		role, exists := c.Get("userRole")
		if !exists {
			utils.RespondError(c, http.StatusUnauthorized, "Role user tidak ditemukan", nil, logger)
			c.Abort()
			return
		}

		userRole, ok := role.(string)
		if !ok {
			utils.RespondError(c, http.StatusUnauthorized, "Role user tidak valid", nil, logger)
			c.Abort()
			return
		}

		// Periksa apakah role user termasuk dalam daftar role yang diizinkan
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				c.Next()
				return
			}
		}

		// Jika role tidak diizinkan, tolak akses
		utils.RespondError(c, http.StatusForbidden, "Akses ditolak", nil, logger)
		c.Abort()
	}
}
