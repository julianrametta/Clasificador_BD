# Dockerfile
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY src/go.mod src/go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY src/ .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest  
WORKDIR /root/

# Install bash and curl
RUN apk --no-cache add bash curl

# Download wait-for-it.sh
RUN curl -o /wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Command to run the executable
CMD ["/wait-for-it.sh", "db_clasificador:3306", "--", "./main"]