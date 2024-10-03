# Use the official Golang image as the base image
FROM golang:1.22.4-bullseye AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all the dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN make && ./scripts/database/setup.sh sqlite

# Use a minimal image for the final build
FROM debian:bullseye-slim

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/mood /app/mood.db ./

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the executable
CMD ["/app/mood", "--addr=:8080"]
