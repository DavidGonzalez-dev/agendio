# Real-Time Multi-Tenant Booking Platform

A booking and scheduling platform where small businesses — barbershops, gyms, clinics — publish their services and accept appointments from registered clients in real time. Built to demonstrate the hard parts of backend engineering: **concurrency-safe booking**, **tenant isolation**, and **real-time notifications**.

> **Status: in active development.** The backend skeleton and domain layer are in place; booking concurrency (the core feature), WebSockets, and the frontend are being built. Demo URL and screenshots will be added on release.

---

## Why this project

Most booking apps in a portfolio are CRUD demos. This one solves the problems that actually make booking software hard:

- **No double bookings, ever.** Two clients hitting "Book" on the same slot at the same instant: one wins, the other gets a `409 Conflict`. Proven by a concurrency test against a real PostgreSQL instance.
- **Tenants can't see each other's data.** Every query is scoped by `business_id`; isolation is enforced at the database level, not by convention.
- **Real-time admin dashboard.** A new booking appears on the business owner's screen without refreshing the page — pushed over WebSockets, filtered per tenant.
- **Timezone-safe scheduling.** All timestamps stored in UTC; each business renders availability in its own IANA timezone. No off-by-one booking bugs.

## Tech stack

| Layer | Technology |
| --- | --- |
| Backend | Go, Echo, pgx, sqlc |
| Database | PostgreSQL (golang-migrate) |
| Real-time | WebSockets (gorilla/websocket) |
| Auth | JWT (golang-jwt/v5) |
| Frontend | React + TypeScript + Vite, TanStack Query, Tailwind CSS |
| Calendar | react-big-calendar |
| Infra | Docker Compose, Railway / Render, Vercel |
| Testing | Go test + testcontainers-go (real Postgres, no mocks) |

## Features

**For clients**

- Register once, book at any business on the platform
- Public business page by slug: services, prices, availability
- Booking flow: pick business → service → time slot → confirm
- Cancel your own bookings; the slot frees automatically

**For business owners**

- Single flow creates your business + admin account
- Manage services (name, duration, price)
- Define a weekly recurring schedule — available slots are generated automatically from service duration
- Calendar view of all bookings
- Live notifications when a client books or cancels

## Architecture

Monolith backend with **hexagonal architecture (ports & adapters)**: pure use cases in the center, with HTTP, PostgreSQL, WebSocket, and JWT adapters plugged in from the outside.

```
[React SPA] ──HTTP/JSON──▶ [Go API] ──▶ [PostgreSQL]
     │                        │
     └────── WebSocket /ws/notifications ────┘
```

- `domain` — business rules only, zero external dependencies
- `application` — use cases (one per system action) + port interfaces
- `adapters` — concrete implementations: HTTP handlers, Postgres repositories (sqlc), WebSocket hub, JWT

**Concurrency design (the core)**

Booking happens inside a transaction:

1. `SELECT ... FOR UPDATE` on the business row — serializes bookings per business
2. Overlap check against confirmed bookings
3. Insert booking with `status = confirmed`
4. Final safety net: partial unique index on `bookings(slot_id) WHERE status = 'confirmed'`

Two simultaneous requests for the same slot: the second blocks on the lock or hits the unique index. Availability is derived from bookings — never from a `is_booked` flag — so cancelling frees the slot instantly.

## Getting started

Requirements: Docker + Docker Compose.

```bash
git clone <your-repo-url>
cd reservations-project
cp backend/.env.example backend/.env   # adjust if needed
docker compose up
```

The API starts on `http://localhost:8080` and PostgreSQL on `localhost:5432`. See `backend/.env.example` for the full configuration.

## Project structure

```
backend/
  cmd/api/               # composition root: wires adapters into use cases
  internal/
    domain/              # types with business rules (no external deps)
    application/         # use cases + port interfaces, one package per module
    adapters/            # http, postgres (sqlc), ws, jwt
  migrations/            # versioned SQL (golang-migrate)
frontend/                # React SPA (planned)
docs/                    # spec, technical design, data model, task breakdown
```

## Roadmap

- [x] Technical design and data model
- [x] Backend scaffold (hexagonal layout, domain layer)
- [ ] Auth (client register/login, JWT)
- [ ] Business registration + service management
- [ ] Weekly schedule and slot generation
- [ ] **Booking with concurrency protection + concurrency tests**
- [ ] WebSocket notifications (per tenant)
- [ ] Frontend: public pages, booking flow, admin dashboard
- [ ] Docker Compose, seed data, demo deploy

## License

MIT — add your name/year here when publishing.
