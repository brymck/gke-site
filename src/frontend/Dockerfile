# Base build image
FROM golang:alpine as serverbuilder

# Install dependencies
RUN apk update && apk add --no-cache ca-certificates git && update-ca-certificates
WORKDIR /src

# Use an uncredentialed user
RUN adduser -D -g '' appuser

# Force the Go compiler to use modules
ENV GO111MODULE=on

# Download Go library dependencies
COPY server/go.mod server/go.sum ./
RUN go mod download

# Build
COPY server/ .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/server

# Build Node
FROM node:8-alpine as uibuilder
WORKDIR /usr/src/app
COPY ui/package.json ui/yarn.lock ./
RUN yarn install
COPY ui .
RUN yarn build

# Base deploy image
FROM scratch
COPY --from=serverbuilder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=serverbuilder /etc/passwd /etc/passwd

USER appuser

# Expose container
WORKDIR /frontend
COPY --from=serverbuilder /bin/server ./server
COPY --from=uibuilder /usr/src/app/build ./ui
EXPOSE 8080
ENTRYPOINT ["/frontend/server"]
