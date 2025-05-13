package models

// Status adalah tipe data berbasis string untuk mendefinisikan status umum
type Status string

// Daftar konstanta status yang dapat digunakan di berbagai model
const (
    StatusActive     Status = "active"       // Status untuk entitas yang aktif
    StatusCancelled  Status = "cancelled"    // Status untuk entitas yang dibatalkan
    StatusCompleted  Status = "completed"    // Status untuk entitas yang selesai
    StatusPending    Status = "pending"      // Status untuk entitas yang menunggu proses
    StatusFailed     Status = "failed"       // Status untuk entitas yang gagal
    StatusInProgress Status = "in_progress" // Status untuk entitas yang sedang dalam proses
)

// IsValidStatus memeriksa apakah status yang diberikan valid
func IsValidStatus(status Status) bool {
    validStatuses := []Status{
        StatusActive,
        StatusCancelled,
        StatusCompleted,
        StatusPending,
        StatusFailed,
        StatusInProgress,
    }

    for _, validStatus := range validStatuses {
        if status == validStatus {
            return true
        }
    }
    return false
}