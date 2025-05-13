package utils

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Status constants
const (
	StatusError   = "error"
	StatusSuccess = "success"
)

// ErrorResponse adalah struktur untuk respons error
type ErrorResponse struct {
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	Details   interface{} `json:"details,omitempty"`
	Timestamp string      `json:"timestamp"`
}

// SuccessResponse adalah struktur untuk respons sukses
type SuccessResponse struct {
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp string      `json:"timestamp"`
}

// Logger adalah interface untuk logging, memungkinkan penggunaan library logging lain
type Logger interface {
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
}

// DefaultLogger adalah implementasi default menggunakan log.Printf
type DefaultLogger struct{}

// Errorf log error dengan prefix [ERROR]
func (l *DefaultLogger) Errorf(format string, args ...interface{}) {
	log.Printf("[ERROR] "+format, args...)
}

// Infof log info dengan prefix [INFO]
func (l *DefaultLogger) Infof(format string, args ...interface{}) {
	log.Printf("[INFO] "+format, args...)
}

// DefaultLoggerInstance adalah instance global yang bisa digunakan di seluruh aplikasi
var DefaultLoggerInstance Logger = &DefaultLogger{}

// RespondError mengirimkan response error dalam format JSON
func RespondError(c *gin.Context, statusCode int, message string, details interface{}, logger Logger) {
	if logger == nil {
		logger = DefaultLoggerInstance
	}

	// Logging error untuk debugging
	logger.Errorf("Error: %s - Details: %v", message, details)

	// Validasi status code
	if statusCode < 400 || statusCode > 599 {
		statusCode = http.StatusInternalServerError
	}

	// Kirimkan respons error
	c.JSON(statusCode, ErrorResponse{
		Status:    StatusError,
		Message:   message,
		Details:   details,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// RespondSuccess mengirimkan response sukses dalam format JSON
func RespondSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	// Validasi status code
	if statusCode < 200 || statusCode > 299 {
		statusCode = http.StatusOK
	}

	// Kirimkan respons sukses
	c.JSON(statusCode, SuccessResponse{
		Status:    StatusSuccess,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}
