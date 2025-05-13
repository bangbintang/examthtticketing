package middleware

import (
    "bytes"
    "encoding/json"
    "io"
    "log"
    "time"

    "github.com/gin-gonic/gin"
)

type LogEntry struct {
    Timestamp    string      `json:"timestamp"`
    ClientIP     string      `json:"client_ip"`
    Method       string      `json:"method"`
    URI          string      `json:"uri"`
    StatusCode   int         `json:"status_code"`
    Duration     string      `json:"duration"`
    RequestBody  interface{} `json:"request_body,omitempty"`
    ResponseSize int         `json:"response_size"`
    Error        string      `json:"error,omitempty"`
}

func LoggingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        startTime := time.Now()

        // Baca body permintaan
        var requestBody interface{}
        if c.Request.Body != nil {
            bodyBytes, err := io.ReadAll(c.Request.Body)
            if err != nil {
                log.Printf("Error membaca body permintaan: %v", err)
            } else {
                // Reset body permintaan agar dapat dibaca ulang oleh handler
                c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

                // Decode JSON jika memungkinkan
                if json.Valid(bodyBytes) {
                    _ = json.Unmarshal(bodyBytes, &requestBody)
                } else {
                    requestBody = string(bodyBytes)
                }
            }
        }

        // Lanjutkan ke handler berikutnya
        c.Next()

        // Hitung durasi permintaan
        duration := time.Since(startTime)

        // Buat log entry
        logEntry := LogEntry{
            Timestamp:    startTime.Format(time.RFC3339),
            ClientIP:     c.ClientIP(),
            Method:       c.Request.Method,
            URI:          c.Request.RequestURI,
            StatusCode:   c.Writer.Status(),
            Duration:     duration.String(),
            RequestBody:  requestBody,
            ResponseSize: c.Writer.Size(),
        }

        // Tambahkan error jika ada
        if len(c.Errors) > 0 {
            logEntry.Error = c.Errors.String()
        }

        // Catat log dalam format JSON
        logData, _ := json.Marshal(logEntry)
        log.Println(string(logData))
    }
}