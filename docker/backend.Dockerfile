# Build
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
WORKDIR /app/backend
RUN go build -o backend

# Run
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/backend/backend .
EXPOSE 8080
CMD ["./backend"]
