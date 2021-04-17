# Start from golang base image
FROM golang:1.13-alpine as builder

# ENV GO111MODULE=on
# ARG
ARG PORT
ARG READER_HOST
ARG READER_PORT
ARG READER_USER
ARG READER_PASSWORD
ARG WRITER_HOST
ARG WRITER_PORT
ARG WRITER_USER
ARG TIMEOUT_ON_SECONDS
ARG OPERATION_ON_EACH_CONTEXT
# ENV
ENV PORT=${PORT}
ENV READER_HOST=${READER_HOST}
ENV READER_PORT=${READER_PORT}
ENV READER_USER=${READER_USER}
ENV READER_PASSWORD=${READER_PASSWORD}
ENV WRITER_HOST=${WRITER_HOST}
ENV WRITER_PORT=${WRITER_PORT}
ENV WRITER_USER=${WRITER_USER}
ENV TIMEOUT_ON_SECONDS=${TIMEOUT_ON_SECONDS}
ENV OPERATION_ON_EACH_CONTEXT=${TIMEOUT_ON_SECONDS}

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
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .       

# Expose port 8800 to the outside world
EXPOSE 8800

#Command to run the executable
CMD ["./main"]