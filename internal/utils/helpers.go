package utils

import (
    "time"

    "github.com/google/uuid"
)

// FormatDate mengubah waktu menjadi string dengan format tertentu.
// Jika format kosong, akan menggunakan format default "02 Jan 2006".
func FormatDate(t time.Time, format string) string {
    if format == "" {
        format = "02 Jan 2006"
    }
    return t.Format(format)
}

// GenerateUUID membuat UUID baru dalam format string.
func GenerateUUID() string {
    return uuid.New().String()
}

// GetCurrentTime mengembalikan waktu sekarang.
// Fungsi ini dapat diubah untuk menerima dependency waktu agar lebih mudah diuji.
func GetCurrentTime() time.Time {
    return time.Now()
}