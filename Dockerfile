<<<<<<< HEAD
FROM golang:1.21


# Set destination for COPY
WORKDIR /APP

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./


# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-app


# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage


WORKDIR /

COPY --from=build-stage /docker-app /docker-app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["docker-app"]


=======
from base 
>>>>>>> main
