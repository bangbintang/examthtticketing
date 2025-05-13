package utils

import (
    "errors"
    "regexp"
)

// IsValidEmail memeriksa apakah email valid menggunakan regex yang lebih komprehensif
func IsValidEmail(email string) bool {
    // Regex sesuai standar RFC 5322
    re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
    return re.MatchString(email)
}

// ValidatePassword memeriksa apakah password memenuhi kriteria minimal
func ValidatePassword(password string) error {
    if len(password) < 8 {
        return errors.New("password harus minimal 8 karakter")
    }

    // Harus mengandung setidaknya satu angka
    if !regexp.MustCompile(`[0-9]`).MatchString(password) {
        return errors.New("password harus mengandung setidaknya satu angka")
    }

    // Harus mengandung setidaknya satu huruf besar
    if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
        return errors.New("password harus mengandung setidaknya satu huruf besar")
    }

    // Harus mengandung setidaknya satu simbol
    if !regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};':",\.<>\/?\\|]`).MatchString(password) {
        return errors.New("password harus mengandung setidaknya satu simbol")
    }

    return nil
}

// IsNotEmpty memeriksa apakah string tidak kosong
func IsNotEmpty(s string) bool {
    return len(s) > 0
}