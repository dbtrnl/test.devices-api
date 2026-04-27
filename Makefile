BIN_FILENAME=devices-api
BIN_PATH=./bin
CMD_PATH=./cmd/api

ENV?=local

# Metadata embedded on compiled binary
VERSION ?= $(shell git describe --tags --always --dirty)
COMMIT  ?= $(shell git rev-parse --short HEAD)
DATE    ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS = -ldflags "\
	-X github.com/dbtrnl/test.devices-api/internal/infra/buildinfo.Version=$(VERSION) \
	-X github.com/dbtrnl/test.devices-api/internal/infra/buildinfo.Commit=$(COMMIT) \
	-X github.com/dbtrnl/test.devices-api/internal/infra/buildinfo.BuildTime=$(DATE)"

help:
	@echo "Available commands:"
	@echo "  build      		- Build the binary for Linux amd64"
	@echo "  clean      		- Remove the bin directory"
	@echo "  clean-db   		- Stop and remove the PostgreSQL container"
	@echo "  init-db    		- Initialize the PostgreSQL database container"
	@echo "  run        		- Clean, build, and run the binary with the specified ENV"
	@echo "  run-local  		- Run the binary in local environment"
	@echo "  test-integration	- Initializes a local DB, compiles the binary and runs the integration tests"
	@echo "  version    		- Display version information for the built binary"
	@echo "  help       		- Show this help message"

.PHONY: build clean clean-db init-db run run-local test-integration version

build: clean
	mkdir -p $(BIN_PATH)
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_PATH)/$(BIN_FILENAME) $(CMD_PATH)

clean:
	rm -rf $(BIN_PATH)

clean-db:
	docker compose down -v

init-db: clean-db
	docker compose up postgres -d

run: clean build
	ENV=$(ENV) $(BIN_PATH)/$(BIN_FILENAME)

run-local:
	$(MAKE) run ENV=local

test-integration:
	./scripts/test-integration.sh

version:
	go version -m $(BIN_PATH)/$(BIN_FILENAME)