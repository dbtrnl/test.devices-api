#!/usr/bin/env bash

set -e

BIN_FILENAME=devices-api
BIN_PATH=./bin
CMD_PATH=./cmd/api

YELLOW="\033[0;33m"
BLUE="\033[0;34m"
DARK_GRAY="\033[0;90m"
RESET="\033[0m"

printf "${YELLOW}Cleaning DB...\n${DARK_GRAY}"
docker compose down -v > /dev/null 2>&1

printf "${YELLOW}Starting DB...\n${DARK_GRAY}"
docker compose up postgres -d > /dev/null 2>&1

printf "${YELLOW}Waiting for DB to be ready for queries...\n${DARK_GRAY}"
until docker exec devices_db psql -U devices_api -d devices_db -c "SELECT 1" > /dev/null 2>&1; do
  sleep 1
done

printf "${YELLOW}Building binary...\n${DARK_GRAY}"
mkdir -p ${BIN_PATH}
go build -o ${BIN_PATH}/${BIN_FILENAME} ${CMD_PATH}

printf "${YELLOW}Starting API...\n${DARK_GRAY}"
ENV=local ${BIN_PATH}/${BIN_FILENAME} &
API_PID=$!

cleanup() {
  printf "${YELLOW}Stopping API...\n${DARK_GRAY}"
  kill ${API_PID} 2>/dev/null || true
}
trap cleanup EXIT

printf "${YELLOW}Waiting for API to be ready...\n${DARK_GRAY}"
until curl -s http://localhost:8080/health > /dev/null; do
  sleep 1
done

printf "${RESET}Running integration tests...\n${BLUE}"
go test -tags=integration ./... -v