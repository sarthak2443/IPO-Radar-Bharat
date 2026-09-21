FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install ca-certificates and git
RUN apk add --no-cache ca-certificates git

# Cache dependencies
COPY go.mod go.sum* ./
RUN go mod download || true

# Copy source code
COPY . .

# Tidy & build
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /bin/api ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /bin/worker ./cmd/worker

# Production image
FROM alpine:3.21

WORKDIR /app
RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /bin/api /app/api
COPY --from=builder /bin/worker /app/worker
COPY --from=builder /app/.env.example /app/.env.example

EXPOSE 8080

CMD ["/app/api"]
