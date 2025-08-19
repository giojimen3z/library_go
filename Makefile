# --- Env configuration ---
ifneq (,$(wildcard .env))
include .env
export
endif

# --- Quality variables ---
COVER_MIN        ?= 0
COVER_FILE       := coverage.out
PKG              := ./...
GOLANGCI_VERSION := v2.4.0

.PHONY: deps tidy vet test cover lint-install lint \
        compose-up compose-logs compose-down db-psql db-reset \
        swag build test-api compose-restart clean-coverage


# ---------- Restart and Clean Environments  ----------
compose-restart:
	docker compose restart

clean-coverage:
	@rm -f $(COVER_FILE)

clean: clean-coverage
	@go clean -testcache
# ---------- Quality ----------
deps:
	go mod download

tidy:
	go mod tidy

vet:
	go vet $(PKG)

test:
	go test -race -shuffle=on -covermode=atomic -coverprofile=$(COVER_FILE) $(PKG)

cover:
	@$(MAKE) -s test
	@total=$$(go tool cover -func=$(COVER_FILE) | awk '/^total:/ {print $$3}' | tr -d '%'); \
	if awk "BEGIN {exit !($$total >= $(COVER_MIN))}"; then \
	  printf "\033[32m✔ Coverage %.2f%% ≥ min %s%%\033[0m\n" $$total "$(COVER_MIN)"; \
	else \
	  printf "\033[31m✘ Coverage %.2f%% < min %s%%\033[0m\n" $$total "$(COVER_MIN)"; \
	  exit 1; \
	fi

# ---------- Quality ----------
lint-install:
	@echo ">> Removing any existing golangci-lint..."
	@rm -f $$(which golangci-lint) || true
	@echo ">> Installing golangci-lint $(GOLANGCI_VERSION)..."
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
	  sh -s -- -b $$(go env GOPATH)/bin $(GOLANGCI_VERSION)

lint: lint-install
	@golangci-lint version
	@golangci-lint run ./...

# ---------- DevOps / Runtime ----------
compose-up:
	docker compose up -d --build

compose-logs:
	docker compose logs -f

compose-down:
	docker compose down -v

db-psql:
	docker exec -it library-db psql -U $(DB_USER) -d $(DB_NAME)

db-reset: compose-down compose-up

swag:
	swag init -g cmd/api/main.go

build:
	docker build -t library-api:local .

test-api:
	bruno run bruno_collection/