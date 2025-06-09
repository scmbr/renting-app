# Build stage
FROM golang:1.23 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd

# Run stage
FROM debian:bullseye-slim
WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
