# Dokumentasi Arsitektur Aplikasi Ticketing Konser

## 1. Overview

Aplikasi Ticketing Konser ini dibangun dengan arsitektur berbasis REST API menggunakan bahasa Go (Golang) dan framework Gin untuk web server. Database yang digunakan adalah PostgreSQL dengan ORM GORM untuk memudahkan interaksi data.

## 2. Struktur Aplikasi

- **cmd/**  
  Folder untuk entry point aplikasi (misal `main.go`).

- **internal/**  
  Folder utama berisi kode aplikasi yang dibagi menjadi beberapa subfolder:
  - **handlers/**: Berisi HTTP handler yang menerima request dan mengirim response.
  - **service/**: Berisi logika bisnis aplikasi, berinteraksi dengan database melalui repository.
  - **utils/**: Berisi utilitas seperti JWT helper, validasi, dll.
  - **middleware/**: Berisi middleware untuk autentikasi, logging, dan RBAC.
  - **models/**: Berisi definisi model data.

## 3. Alur Request

1. Client mengirimkan HTTP request ke server.
2. Request melewati middleware (logging, autentikasi, RBAC).
3. Jika autentikasi dan otorisasi berhasil, request diteruskan ke handler sesuai route.
4. Handler memanggil service untuk menjalankan logika bisnis.
5. Service berinteraksi dengan database melalui ORM GORM.
6. Hasil dari service dikembalikan ke handler, lalu dikirim sebagai response ke client.

## 4. Autentikasi dan Otorisasi

- Menggunakan JWT (JSON Web Token) untuk autentikasi.
- Middleware autentikasi memvalidasi token dan mengekstrak klaim user (userID, role).
- Middleware RBAC membatasi akses berdasarkan role user (misal hanya admin yang bisa mengakses route tertentu).

## 5. Database

- PostgreSQL sebagai database utama.
- GORM sebagai ORM untuk memudahkan operasi CRUD dan migrasi schema.

## 6. Middleware

- **LoggingMiddleware**: Mencatat log request dan response.
- **AuthMiddleware**: Memvalidasi token JWT dan mengatur context user.
- **AdminMiddleware**: Membatasi akses hanya untuk user dengan role admin.

## 7. Deployment

- Aplikasi dapat dijalankan di server lokal atau cloud dengan environment variable untuk konfigurasi (misal `DATABASE_DSN`, `JWT_SECRET`).
- Port default aplikasi adalah 8080, bisa diubah sesuai kebutuhan.

---

Dokumentasi ini bisa dikembangkan lebih detail sesuai kebutuhan, misalnya menambahkan diagram arsitektur atau flowchart.

Apakah Anda ingin saya bantu buatkan diagram arsitektur sederhana untuk aplikasi ini?
