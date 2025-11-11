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
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pushOff ./main.go

# Use a smaller, secure image for the final stage
FROM alpine:3.22.2

# Install openssh-client to get ssh-keygen
RUN apk add --no-cache openssh-client=10.0_p1-r9
# Set the working directory inside the container
WORKDIR /app

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/pushOff .
# Create keys folder
RUN mkdir -p /app/keys && chmod 700 keys
COPY --from=builder /app/authorized_keys /app/keys/.
RUN chmod 600 /app/keys/authorized_keys

# Expose port 23234 to the outside world
EXPOSE 23234

# Command to run the executable
CMD ["./pushOff"]

