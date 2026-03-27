# Use the official Golang image for building
FROM golang:1.26.1-alpine AS builder
# Set working directory
WORKDIR /build
# Copy Go modules and dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy source code
COPY . .
# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .
# Use a minimal base image for final deployment
FROM alpine:latest
# Set working directory in the container
WORKDIR /app
# Add non-root user
RUN adduser -D myuser
USER myuser
# Copy the built binary from the builder stage
COPY --from=builder /build/main .
# Expose the application port
EXPOSE 8080
# Run the application
CMD ["./main"]