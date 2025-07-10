# 🧑‍✈️ Voucher Assignment Frontend (React)

This is the frontend interface for airline crew to generate 3 random seat vouchers per flight, using the voucher backend service.

---

## 🚀 Features

- Enter flight and crew details
- Select aircraft type from dropdown
- Date picker for flight date
- Generate 3 random seat vouchers via backend
- Prevents duplicate voucher generation
- Displays error message if already generated

---

## 🛠 Tech Stack

- React (Vite or Create React App)
- TailwindCSS (optional, used for layout/styling)
- Axios (for API requests)

---

## 📦 Getting Started

### 1. Install dependencies

```bash
npm install
```

### 2. Start the dev server

```bash
npm start
```

Runs the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in your browser.

---

## 🌐 Backend Connection

Make sure your backend is running at:

```
http://localhost:8080
```

If hosted elsewhere, update the URLs inside `src/App.jsx`.

---

## 📤 Example Payload Sent

```json
{
  "name": "Sarah",
  "id": "98123",
  "flightNumber": "ID102",
  "date": "2025-07-12",
  "aircraft": "Airbus 320"
}
```

---

## 🧪 Sample CURL (Backend Reference)

```bash
curl -X POST http://localhost:8080/api/check \
  -H "Content-Type: application/json" \
  -d '{"flightNumber":"GA102", "date":"2025-07-12"}'
```

```bash
curl -X POST http://localhost:8080/api/generate \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Sarah",
    "id": "98123",
    "flightNumber": "ID102",
    "date": "2025-07-12",
    "aircraft": "Airbus 320"
  }'
```

---

## ✅ License

MIT © Iman K
