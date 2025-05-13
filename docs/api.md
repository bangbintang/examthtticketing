# Dokumentasi API Ticketing Konser

## Autentikasi

### POST /auth/login

Login user dan mendapatkan token JWT.

**Request Body:**

```json
{
  "email": "user@example.com",
  "password": "password123"
}

Response:

json

{
  "token": "jwt_token",
  "refresh_token": "refresh_token"
}

POST /auth/refresh
Refresh token JWT.

Request Body:

{
  "refresh_token": "refresh_token"
}

Response:
{
  "token": "new_jwt_token"
}


User
POST /users/register
Registrasi user baru.

Request Body:
{
  "name": "Nama User",
  "email": "user@example.com",
  "password": "password123"
}

Response:
{
  "id": "user_id",
  "name": "Nama User",
  "email": "user@example.com"
}


Ticket
POST /tickets/purchase
Membeli tiket (perlu autentikasi).

Headers:
Authorization: Bearer <token>

Request Body:
{
  "event_id": "event_id",
  "quantity": 2
}

Response:
{
  "ticket_id": "ticket_id",
  "event_id": "event_id",
  "user_id": "user_id",
  "quantity": 2,
  "purchase_date": "2025-05-03T10:00:00Z"
}


GET /tickets/user/:userID
Mendapatkan tiket milik user tertentu (perlu autentikasi).

Headers:
Authorization: Bearer <token>

Response:
[
  {
    "ticket_id": "ticket_id_1",
    "event_id": "event_id_1",
    "quantity": 2,
    "purchase_date": "2025-05-01T10:00:00Z"
  },
  {
    "ticket_id": "ticket_id_2",
    "event_id": "event_id_2",
    "quantity": 1,
    "purchase_date": "2025-05-02T11:00:00Z"
  }
]


Review
POST /reviews
Membuat review baru (perlu autentikasi).

Headers:
Authorization: Bearer <token>

Request Body:
{
  "event_id": "event_id",
  "rating": 4,
  "comment": "Acara sangat menyenangkan!"
}

Response:
{
  "review_id": "review_id",
  "event_id": "event_id",
  "user_id": "user_id",
  "rating": 4,
  "comment": "Acara sangat menyenangkan!",
  "created_at": "2025-05-03T12:00:00Z"
}



GET /reviews/event/:eventID
Mendapatkan review untuk event tertentu (publik).

Response:
[
  {
    "review_id": "review_id_1",
    "user_id": "user_id_1",
    "rating": 5,
    "comment": "Luar biasa!",
    "created_at": "2025-05-01T09:00:00Z"
  },
  {
    "review_id": "review_id_2",
    "user_id": "user_id_2",
    "rating": 3,
    "comment": "Bagus tapi bisa lebih baik.",
    "created_at": "2025-05-02T10:00:00Z"
  }
]


Transaction
POST /transactions
Membuat transaksi baru (perlu autentikasi).

Headers:
Authorization: Bearer <token>

Request Body:
{
  "ticket_id": "ticket_id",
  "amount": 100000
}

Response:
{
  "transaction_id": "transaction_id",
  "ticket_id": "ticket_id",
  "amount": 100000,
  "transaction_date": "2025-05-03T13:00:00Z"
}



GET /transactions/user/:userID
Mendapatkan transaksi milik user tertentu (perlu autentikasi).

Headers:
Authorization: Bearer <token>

Response:
[
  {
    "transaction_id": "transaction_id_1",
    "ticket_id": "ticket_id_1",
    "amount": 100000,
    "transaction_date": "2025-05-01T10:00:00Z"
  },
  {
    "transaction_id": "transaction_id_2",
    "ticket_id": "ticket_id_2",
    "amount": 150000,
    "transaction_date": "2025-05-02T11:00:00Z"
  }
]


Admin (hanya untuk role admin)
POST /admin/events
Membuat event baru.

PUT /admin/events/:id
Mengubah event.

DELETE /admin/events/:id
Menghapus event.
```
