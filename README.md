# Library Go

**Library Go** is a practice project to learn and assess **Golang** using **Hexagonal Architecture (Ports & Adapters)**.  
The domain models a small **library system**: `Book`, `Author`, `Member`, `Loan`, and `Reservation`.

> This repo is **for learning** (not production). It focuses on clean design, testing, and DevEx (Makefile, Docker, CI, Lint).

---

## ✨ Features

- **Go 1.24** with **Hexagonal Architecture** (domain-first).
- REST API (Gin) scaffold under `cmd/api` (health, example endpoints; extend with your use cases).
- **PostgreSQL** via Docker Compose.
- **GORM** (optional) or ports for custom persistence adapters.
- **Swagger (swag)** for API docs.
- **Makefile**: tidy/vet/lint/test/coverage/compose.
- **GitHub Actions**: separated **CI** (tests + coverage) and **Lint** workflows (PR only).
- **golangci-lint** with configurable rules, skipping `config/`, `mocks/`, `builder/`.
- **Bruno** collection for functional API checks (optional).

---

## 🗂 Project Structure

```text
.
├── cmd/
│   └── api/                # Entry point (main.go) and HTTP wiring
├── internal/               # (recommended) app internals not meant as public API
│   ├── application/        # Use cases / services (orchestrates domain ports)
│   ├── domain/             # Entities, value objects, domain ports (interfaces)
│   ├── infrastructure/     # Adapters: DB, HTTP handlers, config, logging
│   └── test/               # Test builders/mocks/utilities
├── migration/              # SQL migrations (if used)
├── pkg/                    # Shared helper packages (public, stable)
├── scripts/                # Dev/ops scripts
├── bruno_collection/       # Bruno API tests (optional)
├── .github/workflows/      # CI and Lint workflows
└── README.md
```

**Layer responsibilities**
- **domain**: business rules, pure Go, no frameworks.
- **application**: use cases; depends on domain ports, not on adapters.
- **infrastructure**: adapters (HTTP, DB, config), maps to/from domain.

---

## ✅ Requirements

- Go **1.24+**
- Docker & Docker Compose
- (Optional) [Bruno](https://www.usebruno.com/) CLI for API tests
- Make (GNU make)

---

## ⚙️ Setup

### 1) Clone & Modules
```bash
git clone https://github.com/giojimen3z/library_go
cd library_go
go mod tidy
```

### 2) Environment
Create a `.env` from this minimal example:
```env
PORT=8080
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=library
DB_SSLMODE=disable
```

> Tip: el repo ignora `*.env`. Sube un `example.env` si quieres referencia.

---

## ▶️ Run (Docker)

```bash
make compose-up     # builds API image and boots Postgres + API
make compose-logs   # follow logs
make compose-down   # stop & remove containers/volumes
```

API:
```
http://localhost:8080
```

Swagger (si generas docs con swag, ver abajo):
```
http://localhost:8080/swagger/index.html
```

---

## 🧪 Test, Coverage & Lint

```bash
make test                 # run unit tests
make COVER_MIN=0 cover    # run tests + coverage gate (default 0 while bootstrapping)
make lint                 # golangci-lint (uses .golangci.yml)
```

- **Coverage gate**: `COVER_MIN` configurable (en CI queda 0 al inicio).
- **Lint config**: ver `.golangci.yml` (excluye `config/`, `mocks/`, `builder/`; habilita `govet`, `staticcheck`, `revive`, etc.).

---

## 🧰 Makefile Targets (main)

```text
deps          # go mod download
tidy          # go mod tidy
vet           # go vet ./...
test          # go test with race+coverprofile
cover         # test + coverage threshold (green/red)
lint          # golangci-lint run (instalará si no está)
build         # docker build API image
compose-up    # docker compose up -d --build
compose-logs  # docker compose logs -f
compose-down  # docker compose down -v
db-psql       # psql inside the db container
swag          # swag init -g cmd/api/main.go (generates docs/)
test-api      # bruno run bruno_collection/
```

---

## 📚 Swagger (OpenAPI)

Generar/actualizar documentación (requiere `swag`):
```bash
make swag
# o
swag init -g cmd/api/main.go -o docs
```

> Nota: si usas una versión de `swag` que introduce campos incompatibles, ajusta el archivo generado o actualiza `swag`. Mantén los annotations en handlers para que los juniors vean el flujo.

---

## 🤖 CI/CD (GitHub Actions)

Workflows separados, **solo en pull_request** (GitFlow: `master`, `develop`, `feature/*`, `release/*`, `hotfix/*`):

- **Library CI**: ejecuta tests y coverage gate.
- **Lint**: ejecuta `golangci-lint` (`v1.64.7`, compatible con Go 1.24).

> Cuando quieras bloquear merges: activa **Branch protection rules** (Require PR, Require 1 approval, Require status checks: *Library CI* y *Lint*).

---

## 🧱 Domain Model (initial sketch)

- **Book**: id, title, isbn, publishedAt, authors[], copiesAvailable
- **Author**: id, name, bio?
- **Member**: id, name, email, status
- **Loan**: id, memberId, bookId, loanDate, dueDate, returnDate?
- **Reservation**: id, memberId, bookId, status, createdAt

> Reglas típicas: disponibilidad de copias, reservas en espera, fechas de préstamo/retorno, multas (futuro), etc.

---

## 🗺 Roadmap (suggested)

- CRUD básico para `Book`, `Author`, `Member`.
- Casos de uso: `CreateLoan`, `ReturnLoan`, `CreateReservation`, `CancelReservation`.
- Persistencia vía puerto (interface) + adapter GORM/Postgres.
- Validaciones de dominio (no en controladores).
- Observabilidad mínima (logger `slog`).
- Tests: domain puro + adapters con SQLMock o DB de test.
- Incrementar `COVER_MIN` gradualmente en CI.

---

## 📜 License

MIT
