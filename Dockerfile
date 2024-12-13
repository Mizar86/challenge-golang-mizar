
FROM golang:1.20 AS builder

WORKDIR /app

# Copy go.mod and go.sum for dependency resolution
COPY go.mod go.sum ./

RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the application binary
RUN go build -o carpooling-service ./cmd/main.go


FROM alpine:3.13

WORKDIR /app

COPY --from=builder /app/carpooling-service .

# Expose the port
EXPOSE 80

CMD ["./carpooling-service"]
