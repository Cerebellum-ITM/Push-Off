# Dockerfile for Push-Off App

# Use the official Golang image as a base image
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pusgOff ./main.go

# Use a smaller, secure image for the final stage
FROM alpine:3.22.2

# Install openssh-client to get ssh-keygen
RUN apk add --no-cache openssh-client=10.2p1

# Set the working directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/main .

# Expose port 23234 to the outside world
EXPOSE 23234

# Command to run the executable
CMD ["./pushOff"]

