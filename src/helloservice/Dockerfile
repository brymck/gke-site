# Base build image
FROM golang:alpine as builder

# Install dependencies
RUN apk update && apk add --no-cache ca-certificates git && update-ca-certificates
WORKDIR /src

# Use an uncredentialed user
RUN adduser -D -g '' appuser

# Force the Go compiler to use modules
ENV GO111MODULE=on

# Download Go library dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/helloservice

# Base deploy image
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

USER appuser

# Expose service on port 3550
COPY --from=builder /bin/helloservice /bin/helloservice
EXPOSE 3550
ENTRYPOINT ["/bin/helloservice"]
