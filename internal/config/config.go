package config

import (
    "encoding/json"
    "errors"
    "log"
    "os"
    "strconv"
    "time"

    "github.com/joho/godotenv"
)

type DatabaseConfig struct {
    Host                 string        `json:"host"`
    Port                 int           `json:"port"`
    Name                 string        `json:"name"`
    Username             string        `json:"username"`
    Password             string        `json:"password"`
    MaxOpenConnection    int           `json:"maxOpenConnection"`
    MaxLifetimeConnection time.Duration `json:"maxLifetimeConnection"` // dalam detik
    MaxIdleConnection    int           `json:"maxIdleConnection"`
    MaxIdleTime          time.Duration `json:"maxIdleTime"`           // dalam detik
}

type Config struct {
    Port                 int            `json:"port"`
    AppName              string         `json:"appName"`
    AppEnv               string         `json:"appEnv"`
    SignatureKey         string         `json:"signatureKey"`
    Database             DatabaseConfig `json:"database"`
    RateLimiterMaxRequest int           `json:"rateLimiterMaxRequest"`
    RateLimiterTimeSecond int           `json:"rateLimiterTimeSecond"`
    JwtSecretKey         string         `json:"jwtSecretKey"`
    JwtExpirationTime    int            `json:"jwtExpirationTime"`
}

// LoadConfig memuat konfigurasi dari file JSON dan environment variables
func LoadConfig(path string) (*Config, error) {
    // Muat environment variables dari file .env jika ada
    if err := godotenv.Load(); err != nil {
        log.Println("Peringatan: file .env tidak ditemukan, menggunakan environment variables default")
    }

    // Baca file konfigurasi JSON
    file, err := os.Open(path)
    if err != nil {
        log.Printf("Error membuka file konfigurasi: %v", err)
        return nil, err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    config := &Config{}
    if err := decoder.Decode(config); err != nil {
        log.Printf("Error mendekode file konfigurasi: %v", err)
        return nil, err
    }

    // Konversi nilai waktu dari detik ke time.Duration (nanosecond)
    config.Database.MaxLifetimeConnection = config.Database.MaxLifetimeConnection * time.Second
    config.Database.MaxIdleTime = config.Database.MaxIdleTime * time.Second

    // Override dengan environment variables jika ada
    overrideWithEnv(config)

    // Validasi konfigurasi
    if err := validateConfig(config); err != nil {
        return nil, err
    }

    return config, nil
}

// overrideWithEnv mengganti nilai konfigurasi dengan environment variables jika tersedia
func overrideWithEnv(config *Config) {
    if port := os.Getenv("APP_PORT"); port != "" {
        if p, err := strconv.Atoi(port); err == nil {
            config.Port = p
        }
    }

    if jwtSecret := os.Getenv("JWT_SECRET_KEY"); jwtSecret != "" {
        config.JwtSecretKey = jwtSecret
    }

    if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
        config.Database.Host = dbHost
    }

    if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
        if p, err := strconv.Atoi(dbPort); err == nil {
            config.Database.Port = p
        }
    }

    // Tambahkan override untuk variabel lain jika diperlukan
}

// validateConfig memvalidasi nilai konfigurasi
func validateConfig(config *Config) error {
    if config.Port < 1 || config.Port > 65535 {
        return errors.New("port harus dalam rentang 1-65535")
    }

    if config.JwtSecretKey == "" {
        return errors.New("JWT secret key tidak boleh kosong")
    }

    if config.Database.Host == "" || config.Database.Name == "" || config.Database.Username == "" {
        return errors.New("konfigurasi database tidak lengkap")
    }

    return nil
}