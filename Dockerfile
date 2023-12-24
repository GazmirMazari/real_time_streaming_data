# First stage: build the Go application
FROM golang:1.21 as build-stage

# Set the working directory in the container
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code from the subdirectory
COPY cmd/svr/*.go ./cmd/svr/

# Build the application
# Assuming your main package is in the cmd/svr directory
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-app ./cmd/svr

# Second stage: create a lean image for running the application
FROM gcr.io/distroless/base-debian11

# Set the working directory in the container
WORKDIR /

# Copy the compiled binary from the build stage
COPY --from=build-stage /docker-app /docker-app

# Expose port 8080
EXPOSE 8080

# Use a non-root user
USER nonroot:nonroot

# Set the binary as the entrypoint
ENTRYPOINT ["/docker-app"]
