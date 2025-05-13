package seed

import (
    "log"
    "ticketing-konser/internal/constants"
    "ticketing-konser/internal/models"

    "gorm.io/gorm"
)

// RunRoleSeeder melakukan seeding data role ke database
func RunRoleSeeder(db *gorm.DB) error {
    // Data role didefinisikan menggunakan konstanta
    roles := []models.Role{
        {
            ID:   1,
            Name: string(constants.Admin), // Menggunakan konstanta dari constants.go
        },
        {
            ID:   2,
            Name: string(constants.Customer), // Menggunakan konstanta dari constants.go
        },
    }

    // Lakukan seeding untuk setiap role
    for _, role := range roles {
        if role.Name == "" {
            log.Printf("Role dengan ID %d memiliki nama kosong, seeding dilewati", role.ID)
            continue
        }

        err := db.FirstOrCreate(&role, models.Role{ID: role.ID}).Error
        if err != nil {
            log.Printf("Gagal melakukan seeding role %d (%s): %v", role.ID, role.Name, err)
            continue // Lanjutkan ke role berikutnya meskipun ada error
        }
        log.Printf("Role %s berhasil disimpan", role.Name)
    }

    return nil
}