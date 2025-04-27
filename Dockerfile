# Use Golang official image
FROM golang:1.24

# Set environment variable
ENV GO111MODULE=on

# Create app directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o app ./cmd/server/main.go

# Expose port
EXPOSE 8080

# Run the app
CMD ["./app"]
