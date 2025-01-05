FROM golang:1.20-alpine

WORKDIR /app

# Install necessary build tools
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Install Swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy source code first
COPY . .

# Generate Swagger documentation
RUN swag init

# Build the application
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the application
CMD ["./main"] 