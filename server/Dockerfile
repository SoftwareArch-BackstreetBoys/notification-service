# Step 1: Build the Go app
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files first to install dependencies
COPY go.mod go.sum ./

# Install Go dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Run go mod tidy to clean up dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o server main.go

# Expose the service port if applicable (e.g., 8080)
EXPOSE 8080

# Command to run the executable
CMD ["./server"]
