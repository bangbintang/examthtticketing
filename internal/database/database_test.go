package database

import (
    "testing"
    "ticketing-konser/internal/models"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func TestAutoMigration(t *testing.T) {
    // Membuka koneksi database SQLite in-memory
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        t.Fatalf("Gagal membuka koneksi database: %v", err)
    }

    // Menjalankan auto migration untuk semua model yang ingin diuji
    err = db.AutoMigrate(
        &models.User{},
        &models.Event{},
        &models.Ticket{},
        &models.Transaction{},
        &models.Review{},
        &models.Notification{},
        &models.AuditTrail{},
        &models.EventReport{},
        &models.ReportSummary{},
    )
    if err != nil {
        t.Fatalf("Auto migration gagal: %v", err)
    }

    // Subtests untuk memeriksa keberadaan tabel
    t.Run("CheckTablesExist", func(t *testing.T) {
        modelsToCheck := []interface{}{
            &models.User{},
            &models.Event{},
            &models.Ticket{},
            &models.Transaction{},
        }

        for _, model := range modelsToCheck {
            t.Run(getModelName(model), func(t *testing.T) {
                if !db.Migrator().HasTable(model) {
                    t.Errorf("Tabel untuk model %T tidak ditemukan setelah migrasi", model)
                }
            })
        }
    })

    // Subtests untuk memeriksa relasi antar model
    t.Run("CheckRelations", func(t *testing.T) {
        // Contoh: Periksa apakah kolom foreign key untuk relasi User -> Ticket ada
        if !db.Migrator().HasColumn(&models.Ticket{}, "user_id") {
            t.Errorf("Kolom foreign key 'user_id' tidak ditemukan di tabel Ticket")
        }
    })

    // Pembersihan database
    t.Cleanup(func() {
        sqlDB, _ := db.DB()
        sqlDB.Close()
    })
}

// getModelName mengembalikan nama model untuk logging
func getModelName(model interface{}) string {
    switch model.(type) {
    case *models.User:
        return "User"
    case *models.Event:
        return "Event"
    case *models.Ticket:
        return "Ticket"
    case *models.Transaction:
        return "Transaction"
    default:
        return "UnknownModel"
    }
}