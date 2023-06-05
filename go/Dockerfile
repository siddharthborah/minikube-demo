# Use the official Golang image as the base
FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the Go application source code to the container
COPY . .

# Build the Go application inside the container
RUN go build -o hello-world-app

# Set the entry point for the container
ENTRYPOINT ["./hello-world-app"]
