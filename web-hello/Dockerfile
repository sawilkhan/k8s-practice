# Stage 1: Build the Go binary
FROM golang:1.22.2-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./

# Download Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Stage 2: Create the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the built Go binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the application will run on
EXPOSE 8080

# Command to run the Go binary
CMD ["./main"]
