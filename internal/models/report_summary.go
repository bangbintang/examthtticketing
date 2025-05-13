package models

import (
    "errors"
    "time"

    "gorm.io/gorm"
)

// ReportType adalah tipe data khusus untuk jenis laporan
type ReportType string

// Nilai-nilai yang valid untuk ReportType
const (
    ReportTypeDaily   ReportType = "daily"
    ReportTypeMonthly ReportType = "monthly"
    ReportTypeYearly  ReportType = "yearly"
)

// ReportSummary adalah model untuk ringkasan laporan
type ReportSummary struct {
    ID                int            `gorm:"primaryKey;autoIncrement"`
    ReportType        ReportType     `gorm:"size:50;not null"` // Gunakan enum untuk ReportType
    PeriodStart       time.Time      `gorm:"type:date;not null;index"` // Tambahkan indeks pada PeriodStart
    PeriodEnd         time.Time      `gorm:"type:date;not null;index"` // Tambahkan indeks pada PeriodEnd
    TotalTicketsSold  int            `gorm:"not null;default:0"`       // Default value untuk TotalTicketsSold
    TotalRevenue      float64        `gorm:"type:decimal(15,2);not null;default:0;check:total_revenue >= 0"` // Validasi nilai positif
    CreatedAt         time.Time      `gorm:"autoCreateTime"`
    UpdatedAt         time.Time      `gorm:"autoUpdateTime"` // Metadata tambahan untuk pembaruan
    DeletedAt         gorm.DeletedAt `gorm:"index"`
}

// BeforeSave adalah hook untuk memvalidasi nilai ReportType sebelum menyimpan data
func (r *ReportSummary) BeforeSave(tx *gorm.DB) (err error) {
    if !isValidReportType(r.ReportType) {
        return errors.New("jenis laporan tidak valid")
    }
    return nil
}

// isValidReportType memeriksa apakah nilai ReportType valid
func isValidReportType(reportType ReportType) bool {
    switch reportType {
    case ReportTypeDaily, ReportTypeMonthly, ReportTypeYearly:
        return true
    default:
        return false
    }
}