# --- Env configuration ---
ifneq (,$(wildcard .env))
include .env
export
endif

# --- Quality variables ---
COVER_MIN        ?= 80
COVER_FILE       := coverage.out
PKG              := ./...
GOLANGCI_VERSION ?= v1.64.7
COVER_EXCLUDE    ?= internal/infrastructure/config|cmd/api|internal/test
FILTERED_COVER_FILE := coverage.filtered.out

# Local tool binaries (do not rely on gvm/GOPATH PATH)
BIN_DIR := $(CURDIR)/.bin
export GOBIN := $(BIN_DIR)
LINT_BIN := $(BIN_DIR)/golangci-lint
GOIMPORTS_BIN := $(BIN_DIR)/goimports
GCI_BIN := $(BIN_DIR)/gci

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
	@# Filtra líneas de archivos a excluir del coverage
	@echo "mode: atomic" > $(FILTERED_COVER_FILE)
	@grep -v -E "$(COVER_EXCLUDE)" $(COVER_FILE) | sed '/^mode:/d' >> $(FILTERED_COVER_FILE)
	@total=$$(go tool cover -func=$(FILTERED_COVER_FILE) | awk '/^total:/ {print $$3}' | tr -d '%'); \
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

#Run Only de DB container
compose-db-up:
	docker compose up -d db

compose-db-logs:
	docker compose logs -f db

compose-db-down:
	docker compose rm -sf db

compose-db-reset:
	docker compose down -v
	docker compose up -d db

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



# -----------------------------
# Tooling: install into .bin/
# -----------------------------

tools:
	@mkdir -p "$(BIN_DIR)"
	@echo ">> Installing goimports into $(BIN_DIR)"
	@GOTOOLCHAIN=$(GO_TOOLCHAIN) GOBIN="$(BIN_DIR)" go install golang.org/x/tools/cmd/goimports@latest
	@echo ">> Installing gci into $(BIN_DIR)"
	@GOTOOLCHAIN=$(GO_TOOLCHAIN) GOBIN="$(BIN_DIR)" go install github.com/daixiang0/gci@v0.13.3

tools-clean:
	@echo ">> Removing $(BIN_DIR)"
	@rm -rf "$(BIN_DIR)"


# -----------------------------
# Formatting: gofmt + goimports + gci
# -----------------------------

fmt: tools
	@echo ">> gofmt -s"
	@gofmt -s -w .
	@echo ">> goimports (local=library)"
	@"$(GOIMPORTS_BIN)" -local library -w .
	@echo ">> gci: std | third-party | local(library)"
	@"$(GCI_BIN)" write -s standard -s default -s 'prefix(library)' --skip-generated .

fmt-check:
	@echo ">> gofmt check"
	@out=$$(gofmt -s -l .); \
	if [ -n "$$out" ]; then \
	  printf "\033[31m✘ gofmt found issues. Run 'make fmt'.\033[0m\n"; \
	  echo "$$out"; exit 1; \
	else \
	  printf "\033[32m✔ gofmt OK\033[0m\n"; \
	fi
	@echo ">> goimports check (local=library)"
	@out=$$("$$(printf %q "$(GOIMPORTS_BIN)")" -local library -l .); \
	if [ -n "$$out" ]; then \
	  printf "\033[31m✘ goimports found issues. Run 'make fmt'.\033[0m\n"; \
	  echo "$$out"; exit 1; \
	else \
	  printf "\033[32m✔ goimports OK\033[0m\n"; \
	fi
	@echo ">> gci check"
	@TMP=$$(mktemp); \
	"$(GCI_BIN)" diff -s standard -s default -s 'prefix(library)' --skip-generated . > $$TMP; \
	if [ -s $$TMP ]; then \
	  printf "\033[31m✘ Import order invalid. Run 'make fmt'.\033[0m\n"; \
	  cat $$TMP; rm -f $$TMP; exit 1; \
	else \
	  printf "\033[32m✔ Imports are OK (order + local=library)\033[0m\n"; \
	  rm -f $$TMP; \
	fi
