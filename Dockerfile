# Use official Golang base image
FROM golang:1.19-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o port-sync src/cmd/app/main.go

# Use a minimal base image for the final stage
FROM alpine:latest
WORKDIR /root/

# Copy the compiled binary
COPY --from=builder /app/port-sync .

# Copy configuration and data files
COPY config.toml ./
COPY data/ ./data/

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./port-sync"]
