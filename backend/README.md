# ✈️ Voucher Seat Assignment Backend (Go + Gin + GORM + MySQL)

This backend service allows airline crew to generate 3 random seat vouchers per flight, based on aircraft type. Built with Go, Gin, GORM, and MySQL.

---

## 🚀 Features

- Generate 3 random unique seat vouchers per flight/date
- Prevent duplicates for same flight & date
- Supports aircraft types: ATR, Airbus 320, Boeing 737 Max
- REST API with JSON input/output
- Uses MySQL (not SQLite)

---

## 🛠 Tech Stack

- Go (Golang)
- Gin (HTTP framework)
- GORM (ORM for MySQL)
- MySQL (via Docker)
- Docker & docker-compose for local dev

---

## 📦 Running the Service

### 1. Clone this repo

```bash
git clone https://github.com/your-org/voucher-backend.git
cd voucher-backend
```

### 2. Build & Start with Docker

```bash
docker-compose up --build
```

> The app will run on: `http://localhost:8080`

---

## 🔌 API Endpoints

### ✅ Check if voucher exists

**POST** `/api/check`

#### Request:

```json
{
  "flightNumber": "GA102",
  "date": "2025-07-12"
}
```

#### Response:

```json
{
  "exists": true
}
```

---

### 🎟 Generate voucher

**POST** `/api/generate`

#### Request:

```json
{
  "name": "Sarah",
  "id": "98123",
  "flightNumber": "ID102",
  "date": "2025-07-12",
  "aircraft": "Airbus 320"
}
```

#### Response:

```json
{
  "success": true,
  "seats": ["3B", "7C", "14D"]
}
```

---

## 🧪 Sample CURL Commands

### Check if voucher exists

```bash
curl -X POST http://localhost:8080/api/check   -H "Content-Type: application/json"   -d '{"flightNumber":"GA102", "date":"2025-07-12"}'
```

### Generate voucher

```bash
curl -X POST http://localhost:8080/api/generate   -H "Content-Type: application/json"   -d '{
    "name": "Sarah",
    "id": "98123",
    "flightNumber": "ID102",
    "date": "2025-07-12",
    "aircraft": "Airbus 320"
  }'
```

---

## 🗃 Database Schema (MySQL)

| Field          | Type     |
|----------------|----------|
| id             | INT PK   |
| crew_name      | TEXT     |
| crew_id        | TEXT     |
| flight_number  | TEXT     |
| flight_date    | TEXT     |
| aircraft_type  | TEXT     |
| seat1          | TEXT     |
| seat2          | TEXT     |
| seat3          | TEXT     |
| created_at     | DATETIME |

---

## ✅ License

MIT © Iman K
