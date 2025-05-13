peran masing-masing direktori dan bagaimana mereka berinteraksi dalam proses bisnis aplikasi 

1. Model
Direktori: models (atau repository jika model disimpan di sana)

Fungsi:
- Model adalah representasi data yang digunakan dalam aplikasi, biasanya berupa struct yang mencerminkan tabel database atau entitas bisnis.
- Model digunakan untuk menyimpan, memproses, dan memvalidasi data.

Proses Bisnis:
- Model digunakan oleh repository untuk berinteraksi dengan database.
- Model juga digunakan oleh service untuk memproses data sebelum dikirim ke handler.


2. Repository
Direktori: repository

Fungsi:
-  Repository bertanggung jawab untuk berinteraksi langsung dengan database.
-  Semua query database (seperti SELECT, INSERT, UPDATE, DELETE) dilakukan di sini.
-  Repository mengembalikan data dalam bentuk model ke service.

Proses Bisnis:
- Repository menerima permintaan dari service untuk mengambil atau memodifikasi data di database.
- Repository mengembalikan hasil query ke service dalam bentuk model.


3. Service
Direktori: service

Fungsi:
- Service bertanggung jawab untuk memproses logika bisnis aplikasi.
- Service menerima data dari handler, memprosesnya (misalnya, validasi, transformasi), dan berinteraksi dengan repository untuk mengambil atau menyimpan data.
- Service mengembalikan hasil ke handler.

Proses Bisnis:
- Service menerima permintaan dari handler.
- Service memproses logika bisnis (misalnya, validasi data, hashing password, dll.).
- Service berinteraksi dengan repository untuk mengambil atau menyimpan data.
- Service mengembalikan hasil ke handler.


4. Handlers
Direktori: handlers

Fungsi:
- Handler bertanggung jawab untuk menangani permintaan HTTP dari klien.
- Handler memanggil service untuk memproses logika bisnis.
- Handler mengembalikan respons HTTP ke klien.

Proses Bisnis:
- Handler menerima permintaan HTTP dari klien.
- Handler memanggil service untuk memproses logika bisnis.
- Handler mengembalikan respons HTTP ke klien.


5. Middleware
Direktori: middleware

Fungsi:
- Middleware bertanggung jawab untuk menangani logika lintas fungsi, seperti autentikasi, logging, atau validasi permintaan.
- Middleware dijalankan sebelum atau setelah handler dipanggil.

Proses Bisnis:
- Middleware dijalankan sebelum handler untuk memvalidasi permintaan (misalnya, autentikasi).
- Middleware dapat menghentikan permintaan jika validasi gagal.


6. Utils
Direktori: utils

Fungsi:
- Utils berisi fungsi-fungsi utilitas yang dapat digunakan di seluruh aplikasi, seperti validasi, format tanggal, atau pembuatan token JWT.

Proses Bisnis:
- Utils digunakan oleh service, repository, atau handler untuk membantu memproses data.