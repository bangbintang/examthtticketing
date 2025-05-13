package seed

import (
    "log"
    "os"
    "ticketing-konser/internal/constants"
    "ticketing-konser/internal/models"

    "github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

// RunUserSeeder melakukan seeding data user (contoh: admin)
func RunUserSeeder(db *gorm.DB) error {
    // Ambil data admin dari environment variables atau gunakan nilai default
    adminName := os.Getenv("ADMIN_NAME")
    if adminName == "" {
        adminName = "Administrator"
    }

    adminEmail := os.Getenv("ADMIN_EMAIL")
    if adminEmail == "" {
        adminEmail = "admin@gmail.com"
    }

    adminPassword := os.Getenv("ADMIN_PASSWORD")
    if adminPassword == "" {
        adminPassword = "admin123"
    }

    // Hash password admin
    password, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("Gagal meng-hash password: %v", err)
        return err
    }

    // Validasi apakah role admin ada di database
    var adminRole models.Role
    if err := db.First(&adminRole, "name = ?", string(constants.Admin)).Error; err != nil {
        log.Printf("Role admin tidak ditemukan: %v", err)
        return err
    }

    // Buat data user admin
    user := models.User{
        ID:       uuid.New(),
        Name:     adminName,
        Email:    adminEmail,
        Password: string(password),
        RoleID:   adminRole.ID, // Pastikan role ID diambil dari database
    }

    // Cari user berdasarkan email, jika tidak ada buat baru
    err = db.FirstOrCreate(&user, models.User{Email: user.Email}).Error
    if err != nil {
        log.Printf("Gagal melakukan seeding user: %v", err)
        return err
    }
    log.Printf("User admin dengan email %s berhasil disimpan", user.Email)

    return nil
}