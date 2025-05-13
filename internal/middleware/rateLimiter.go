package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// RateLimiter adalah struktur untuk menyimpan limiter per IP
type RateLimiter struct {
	limiters map[string]*rate.Limiter
	mu       sync.Mutex
	r        rate.Limit
	b        int
}

// NewRateLimiter membuat instance RateLimiter dengan konfigurasi rate dan burst
func NewRateLimiter(r rate.Limit, b int) *RateLimiter {
	return &RateLimiter{
		limiters: make(map[string]*rate.Limiter),
		r:        r,
		b:        b,
	}
}

// GetLimiter mendapatkan limiter untuk IP tertentu
func (r *RateLimiter) GetLimiter(ip string) *rate.Limiter {
	r.mu.Lock()
	defer r.mu.Unlock()

	limiter, exists := r.limiters[ip]
	if !exists {
		limiter = rate.NewLimiter(r.r, r.b)
		r.limiters[ip] = limiter
	}

	return limiter
}

// RateLimitMiddleware membuat middleware untuk rate limiting
func (r *RateLimiter) RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := r.GetLimiter(ip)

		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Terlalu banyak permintaan, coba lagi nanti",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
