# Use the official Go image as a base
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files first (for dependency caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY main.go .

# Build the Go application
RUN go build -o ping-pong .

# Expose the application's port
EXPOSE 8081

# Command to run the application
CMD ["./ping-pong"]
