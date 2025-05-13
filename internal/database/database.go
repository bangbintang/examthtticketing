package database

import (
    "fmt"
    "log"
    "os"
    "time"
    "ticketing-konser/internal/config"
    "ticketing-konser/internal/database/seed"
    "ticketing-konser/internal/models"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB menginisialisasi koneksi database
func InitDB(cfg *config.Config) (*gorm.DB, error) {
    // Gunakan DSN untuk koneksi database
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
        os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
    )

    gormConfig := &gorm.Config{
        Logger: logger.Default.LogMode(logger.Silent),
    }

    // Tambahkan mekanisme retry untuk koneksi database
    var db *gorm.DB
    var err error
    for i := 0; i < 3; i++ { // Retry hingga 3 kali
        db, err = gorm.Open(postgres.Open(dsn), gormConfig)
        if err == nil {
            break
        }
        log.Printf("Gagal membuka koneksi database, percobaan ke-%d: %v", i+1, err)
        time.Sleep(2 * time.Second)
    }

    if err != nil {
        return nil, fmt.Errorf("gagal membuka koneksi database setelah 3 percobaan: %w", err)
    }

    // Konfigurasi koneksi database
    sqlDB, err := db.DB()
    if err != nil {
        return nil, fmt.Errorf("gagal mendapatkan sql.DB: %w", err)
    }

    sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConnection)
    sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConnection)
    sqlDB.SetConnMaxLifetime(cfg.Database.MaxLifetimeConnection)
    sqlDB.SetConnMaxIdleTime(cfg.Database.MaxIdleTime)

    // Uji koneksi database
    if err := sqlDB.Ping(); err != nil {
        return nil, fmt.Errorf("gagal menguji koneksi database: %w", err)
    }

    DB = db
    log.Println("Database connected successfully")
    return db, nil
}

// CloseDB menutup koneksi database
func CloseDB(db *gorm.DB) {
    sqlDB, err := db.DB()
    if err != nil {
        log.Printf("Gagal mendapatkan sql.DB: %v", err)
        return
    }
    if err := sqlDB.Close(); err != nil {
        log.Printf("Gagal menutup koneksi database: %v", err)
    }
}

// MigrateDB menjalankan migrasi database
func MigrateDB(db *gorm.DB) error {
    log.Println("Memulai migrasi database...")

    // Daftar model yang akan dimigrasikan
    modelsToMigrate := []interface{}{
        &models.Role{},
        &models.User{},
        // &models.Event{},
        // &models.Ticket{},
        // &models.Transaction{},
        // &models.AuditTrail{},
        // &models.ReportSummary{},
        // &models.EventReport{},
        // &models.Notification{},
        // &models.Review{},
    }

    // Jalankan migrasi untuk setiap model
    for _, model := range modelsToMigrate {
        log.Printf("Migrasi model: %T", model)
        if err := db.AutoMigrate(model); err != nil {
            return fmt.Errorf("gagal migrasi model %T: %w", model, err)
        }
    }

    log.Println("Database migration completed successfully")
    return nil
}

// SeedDB menjalankan seeding data ke database
func SeedDB(db *gorm.DB) error {
    log.Println("Memulai seeding database...")

    // Periksa apakah tabel Role sudah memiliki data
    var roleCount int64
    if err := db.Model(&models.Role{}).Count(&roleCount).Error; err != nil {
        return fmt.Errorf("gagal memeriksa data role: %w", err)
    }
    if roleCount == 0 {
        log.Println("Seeding data untuk tabel Role...")
        if err := seed.RunRoleSeeder(db); err != nil {
            return fmt.Errorf("gagal melakukan seeding role: %w", err)
        }
    } else {
        log.Println("Data role sudah ada, melewati seeding role")
    }

    // Periksa apakah tabel User sudah memiliki data
    var userCount int64
    if err := db.Model(&models.User{}).Count(&userCount).Error; err != nil {
        return fmt.Errorf("gagal memeriksa data user: %w", err)
    }
    if userCount == 0 {
        log.Println("Seeding data untuk tabel User...")
        if err := seed.RunUserSeeder(db); err != nil {
            return fmt.Errorf("gagal melakukan seeding user: %w", err)
        }
    } else {
        log.Println("Data user sudah ada, melewati seeding user")
    }

    log.Println("Database seeding completed successfully")
    return nil
}