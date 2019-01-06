FROM golang:1.11.2-alpine3.8 as builder
RUN apk add --no-cache ca-certificates git curl
RUN curl -L -s -o /go/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64
RUN chmod +x /go/bin/dep

ENV PROJECT github.com/brymck/gke-site/src/helloservice
WORKDIR /go/src/$PROJECT

# restore dependencies
COPY Gopkg.* ./
RUN dep ensure --vendor-only -v

COPY . .
RUN go build -o /helloservice .

FROM alpine:3.8 AS release
RUN apk add --no-cache ca-certificates
# RUN GRPC_HEALTH_PROBE_VERSION=v0.2.0 && \
#     wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
#     chmod +x /bin/grpc_health_probe
WORKDIR /helloservice
COPY --from=builder /helloservice ./server
EXPOSE 3550
ENTRYPOINT ["/helloservice/server"]
