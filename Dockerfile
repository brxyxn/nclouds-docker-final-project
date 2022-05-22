# Start from golang base image
FROM golang:1.18.2-alpine3.15 as builder

ENV GO111MODULE=on
ENV ENV=Development

# Add Maintainer info
LABEL maintainer="Brayan Lopez <brxyxn.corp@live.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY ./backend/ ./backend

# Build the Go app
RUN go build ./backend/cmd/main.go

#===============================

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

ENV ENV=Production \
    PORT=5000

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .

# Expose port 3000 to the outside world
EXPOSE ${PORT}

# Command to run the executable
CMD ["./main"]
