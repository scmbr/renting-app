
FROM golang:1.23 AS builder

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd

FROM debian:bookworm-slim

WORKDIR /app

# Устанавливаем сертификаты для TLS
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates

EXPOSE 8000

CMD ["./main"]
