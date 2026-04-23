FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o devices-api ./cmd/api

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/devices-api .

EXPOSE 8080

CMD ["./devices-api"]