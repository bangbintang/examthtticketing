package constants

// RoleType adalah tipe data khusus untuk peran pengguna
type RoleType string

const (
    // Admin adalah peran untuk administrator
    Admin RoleType = "admin"
    // Customer adalah peran untuk pelanggan
    Customer RoleType = "customer"
)

// RoleMetadata menyimpan metadata untuk setiap peran
var RoleMetadata = map[RoleType]struct {
    Description string
    Permissions []string
}{
    Admin: {
        Description: "Administrator dengan akses penuh ke sistem",
        Permissions: []string{"manage_users", "manage_events", "view_reports"},
    },
    Customer: {
        Description: "Pelanggan yang dapat membeli tiket dan memberikan ulasan",
        Permissions: []string{"buy_tickets", "write_reviews"},
    },
}