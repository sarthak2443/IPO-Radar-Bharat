# IPO Radar Bharat 🚀

A production-style IPO tracking and analytics platform for India, built as a serious Go/backend learning project.

The goal is to build this progressively — starting with a simple Go API and PostgreSQL backend, then introducing real IPO data ingestion, Kafka, Redis, WebSockets, Docker, Kubernetes, monitoring, and notifications.

---

## 🎯 Project Goals

IPO Radar Bharat provides:

- Upcoming IPOs
- Active IPOs
- Closed/recent IPOs
- IPO details
- Price band and lot size
- Issue size
- Subscription data
- Retail / NII / QIB subscription
- Subscription history
- GMP history (clearly labeled unofficial)
- Listing price and listing performance
- SEBI filing information
- IPO watchlist
- Real-time subscription updates
- Notifications and alerts
- Historical IPO analytics

> **Important:** GMP is not an official exchange/SEBI metric. It should always be clearly labeled as unofficial and the source should be stored.

---

# 🏗️ High-Level Architecture

```text
                         ┌─────────────────────┐
                         │     NSE / BSE       │
                         │   IPO Data Sources  │
                         └──────────┬──────────┘
                                    │
                         ┌──────────▼──────────┐
                         │     Go Ingestion    │
                         │       Workers       │
                         └──────────┬──────────┘
                                    │
                              Kafka Events
                                    │
              ┌─────────────────────┼────────────────────┐
              │                     │                    │
              ▼                     ▼                    ▼
       IPO Processor          Subscription         GMP Processor
              │                     │                    │
              └─────────────────────┼────────────────────┘
                                    │
                                    ▼
                              PostgreSQL
                                    │
                          ┌─────────┴─────────┐
                          │                   │
                       Redis               Go API
                                             │
                                      REST / WebSocket
                                             │
                                             ▼
                                     React Frontend
```

---

# 🧱 Technology Stack

## Frontend (`frontend/`)

- React 19
- TypeScript
- Vite
- Tailwind CSS
- WebSockets

## Backend (`backend/`)

- Go 1.24+
- Echo v4
- GORM
- PostgreSQL
- Redis
- Apache Kafka
- WebSockets

## Infrastructure

- Docker & Docker Compose
- Kubernetes
- Ingress
- Prometheus & Grafana

---

# 📁 Repository Structure

```text
IPO-Radar-Bharat/
│
├── frontend/                   # React + Vite frontend application
│   ├── src/
│   ├── index.html
│   ├── package.json
│   └── ...
│
├── backend/                    # Go API & worker services
│   ├── cmd/
│   │   ├── api/                # API HTTP server entry point
│   │   │   └── main.go
│   │   └── worker/             # Ingestion worker entry point
│   │       └── main.go
│   │
│   ├── internal/
│   │   ├── api/
│   │   │   ├── handlers/       # HTTP route handlers (IPO, Subscription, GMP, Listing)
│   │   │   ├── middleware/     # CORS, Logging, Error handling
│   │   │   └── routes/         # Echo router registration
│   │   │
│   │   ├── domain/             # Core business models & entities
│   │   │   ├── ipo.go
│   │   │   ├── subscription.go
│   │   │   ├── gmp.go
│   │   │   └── listing.go
│   │   │
│   │   ├── repository/         # Database persistence layer (GORM / PostgreSQL)
│   │   ├── service/            # Business logic layer
│   │   ├── ingestion/          # Data ingestion provider interfaces (NSE, BSE, SEBI)
│   │   ├── kafka/              # Kafka event producers & consumers
│   │   ├── cache/              # Redis caching layer
│   │   └── config/             # Environment configuration
│   │
│   ├── migrations/             # SQL database migrations
│   ├── Dockerfile
│   ├── Makefile
│   ├── go.mod
│   └── go.sum
│
├── docker-compose.yml          # Local dev PostgreSQL, Redis
└── README.md
```

---

# 🚀 Quickstart

### Prerequisites

- Go 1.24+
- Node.js 20+ / Bun
- Docker & Docker Compose

### 1. Start Infrastructure (Postgres & Redis)

```bash
docker compose up -d
```

### 2. Run Backend API

```bash
cd backend
cp .env.example .env
make run
# or: go run ./cmd/api
```

The API will start at `http://localhost:8080`.
- Health check: `http://localhost:8080/healthz`
- IPO list: `http://localhost:8080/api/v1/ipos`

### 3. Run Frontend

```bash
cd frontend
bun install   # or: npm install
bun run dev   # or: npm run dev
```

---

# 🚀 Milestone Plan

## Milestone 1 (Current)

- [x] Move frontend into `frontend/`
- [x] Create `backend/`
- [x] Initialize Go module
- [x] Setup Echo v4
- [x] Setup PostgreSQL & GORM
- [x] Create IPO, Subscription, GMP domain models
- [x] Create initial database schema migration
- [x] Build REST API skeleton with routes & handlers
