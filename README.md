# Library Go

**Library Go** is a **practice project** designed to evaluate and improve Golang skills.  
It follows the **Hexagonal Architecture** (Ports and Adapters) to promote clean, maintainable, and testable code.

---

## 📌 Purpose

The main goals of this project are:
- To assess knowledge and skills in **Golang**.
- To practice implementing the **Hexagonal Architecture**.
- To simulate real-world backend development scenarios, including domain-driven design, dependency injection, and testing.

This project is **not** intended for production use.  
It is a **sandbox** for experimenting with concepts and patterns commonly used in professional Go projects.

---

## 🗂 Project Structure

The project follows a layered architecture:

```
cmd/
  api/
    app/
      application/     # Application services (use cases)
      domain/          # Entities, value objects, and domain interfaces
      infrastructure/  # Adapters: DB, HTTP, configuration, controllers
      test/             # Builders, mocks, and test utilities
migration/              # Database migration files
pkg/                    # Shared utilities
scripts/                # Bash/utility scripts
```

**Layer responsibilities:**
- **Domain** → Business rules, entities, and interfaces.
- **Application** → Orchestrates use cases, connects domain logic to infrastructure.
- **Infrastructure** → Implements adapters for persistence, APIs, and delivery mechanisms.
- **Test** → Utilities and mocks for testing.

---

## 🛠 Technologies Used

- **Go 1.22+**
- **Hexagonal Architecture**
- Docker (optional)
- Makefile for automation

---

## 📖 Getting Started

```bash
# Clone the repository
git clone https://github.com/your-username/library_go.git

cd library_go

# Initialize modules
go mod tidy

# Run the application
go run ./cmd/api
```

---

## 🧪 Evaluation Criteria

If used for assessment, the following aspects will be reviewed:
- Correct use of Hexagonal Architecture principles.
- Clean, idiomatic Go code.
- Unit tests and test coverage.
- Proper separation of concerns.
- Error handling and logging.

---

## 📜 License

This project is open source and available under the **MIT License**.
