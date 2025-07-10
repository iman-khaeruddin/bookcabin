# Use an official Go base image
FROM golang:1.23-alpine

# Set working directory
WORKDIR /app

# Install necessary packages
RUN apk add --no-cache git

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port your app listens on
EXPOSE 8080

# Run the binary
CMD ["./main"]

