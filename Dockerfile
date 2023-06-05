# Use the official Golang image as the base
FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the Go application source code to the container
COPY ./go /app/go

# Build the Go application inside the container
WORKDIR /app/go
RUN go build -o hello-world-app

# Set the entry point for the container
ENTRYPOINT ["./hello-world-app"]
