# Use the latest official base image of Go with Alpine Linux
FROM golang:1.22-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files (if applicable)
COPY go.mod go.sum ./

# Update and install system packages
RUN apk update && \
    apk upgrade && \
    apk add --no-cache bash git openssh

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o /message_api

# Remove unnecessary packages and clean up cache to reduce image size
RUN apk del bash git openssh && \
    rm -rf /var/cache/apk/*

# Use a non-root user for running the application
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/message_api"]
