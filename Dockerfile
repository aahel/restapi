# Start from the latest golang base image
FROM golang:alpine AS builder

# Add Maintainer Info
LABEL maintainer="Aahel <aahel1guha@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /build/app

# Copy Go Modules dependency requirements file
COPY go.mod .

# Copy Go Modules expected hashes file
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy all the app sources (recursively copies files and directories from the host into the image)
COPY . .

# Set http port
ENV SERVER_PORT 8081

# Build the app
RUN go build -o restapi

FROM alpine

WORKDIR /app

COPY --from=builder /build/app/config.yaml  .
COPY --from=builder /build/app/swagger.yaml  .
COPY --from=builder /build/app/restapi  .

# Run the app
CMD ["./restapi"]
