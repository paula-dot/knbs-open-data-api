# KNBS Open Data API

An open-source platform for accessing, exploring, and visualizing Kenyan national statistics.

This project provides:

* **A Go-powered backend** exposing structured KNBS datasets through clean REST APIs.
* **A modern React + Vite frontend** for browsing datasets, viewing metadata, and interacting with API documentation.
* **A modular architecture** suitable for future analytics, GIS layers, dashboards, and open-data exploration.

---

## Tech Stack

### **Backend (Go)**

* **Chi** – Fast, lightweight router with middleware support.
* **sqlc** – Type-safe query generation from raw SQL.
* **PostgreSQL** – Reliable, scalable relational database.
* **Goose** – Database migrations.
* **Viper** – Powerful configuration manager.
* **JWT Auth** – Secure, stateless authorization.
* **Clean Architecture** – Clear boundaries between layers.

### **Frontend Features**

* **React + Vite** – Fast dev environment & build system.
* **Tailwind CSS** – Clean, utility-first styling.
* **shadcn/ui** – Polished, production-ready UI components.
* **Zustand** – Lightweight global state management.
* **React Router** – Client-side routing.
<!-- * **Framer Motion** – Smooth animations. -->

---

## Project Structure

```plaintext
knbs-open-api/
│
├── backend/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   ├── config/
│   │   ├── database/
│   │   ├── middleware/
│   │   ├── models/
│   │   ├── handlers/
│   │   ├── repositories/
│   │   └── services/
│   ├── migrations/
│   ├── sqlc/
│   ├── Dockerfile
│   └── go.mod
│
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   ├── pages/
│   │   ├── store/
│   │   ├── lib/
│   │   └── App.jsx
│   ├── public/
│   ├── index.html
│   ├── tailwind.config.js
│   ├── package.json
│   └── Dockerfile
│
└── docker-compose.yml
```

---

## Features

### **Backend**

* Versioned API: `/api/v1/...`
* Dataset listing and retrieval endpoints
* PostgreSQL-backed storage
* Middleware: CORS, logging, recovery, rate limiting
* JWT authentication ready

### **Frontend**

* Dataset listing UI
* Clean navigation: Home, Datasets, Docs
* Tailwind + shadcn UI foundation
* Ready for charts, maps, and analytics modules

---

## Running the Project

### **Requirements**

* Go 1.22+
* Node.js 18+
* PostgreSQL 16+
* Docker (optional)

---

## Backend Setup

```bash
cd backend
go mod tidy
go run cmd/server/main.go
```

Run migrations:

```bash
goose up
```

Generate sqlc code:

```bash
sqlc generate
```

Backend URL:

```plaintext
http://localhost:8080
```

---

## Frontend Setup

```bash
cd frontend
npm install
npm run dev
```

Frontend URL:

```plaintext
http://localhost:5173
```

---

## Docker (Full Stack)

```bash
docker-compose up --build
```

This launches:

* Go backend
* React frontend
* PostgreSQL instance

---

## Roadmap

### Phase 1 — API + Basic UI

* Dataset ingestion
* REST API
* Minimal dataset browsing UI

### Phase 2 — Enhancements

* Search, pagination, filters
* OpenAPI documentation
* Admin dashboard
* Charting + analytics
* GIS (counties, population layers)

### Phase 3 — Public Release

* User accounts
* Custom dashboards
* Dataset exports (CSV/XLSX/JSON)
* KNBS-inspired design system

---

## Contributing

PRs are welcome. Open an issue before major changes.
