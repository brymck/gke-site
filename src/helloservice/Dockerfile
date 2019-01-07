# Base build image
FROM golang:1.11.2-alpine3.8 as builder

# Install dependencies
RUN apk add --no-cache ca-certificates gcc git libc-dev
ENV PROJECT github.com/brymck/gke-site/src/helloservice
WORKDIR /go/src/$PROJECT

# Force the Go compiler to use modules
ENV GO111MODULE=on

# Download Go library dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . .
RUN go build -o /go/bin/helloservice .

# Base deploy image
FROM alpine:3.8 AS release

# Expose service on port 3550
COPY --from=builder /go/bin/helloservice /bin/helloservice
EXPOSE 3550
ENTRYPOINT ["/bin/helloservice"]
