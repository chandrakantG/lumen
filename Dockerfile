# Dockerfile for lumen service
# This Dockerfile builds the lumen service using Go and Alpine Linux.
FROM golang:alpine3.20 AS builder

# Set the environment variables for Go
ENV CGO_ENABLED=0 GOOS=linux 

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# download the dependencies
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Build the application
RUN go build -o lumen  ./cmd/app/main.go

# Final stage
FROM alpine:3.20

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/lumen .

# expose the port the app runs on
EXPOSE 8081

# Set the command to run the application
CMD ["/app/lumen"]