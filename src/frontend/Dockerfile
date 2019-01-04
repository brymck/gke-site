FROM golang:1.11.2-alpine3.8 as serverbuilder
RUN apk add --no-cache ca-certificates git && \
      wget -qO/go/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && \
      chmod +x /go/bin/dep
ENV PROJECT github.com/brymck/gke-site/src/frontend/server
WORKDIR /go/src/$PROJECT

# restore dependencies
COPY server/Gopkg.* ./
RUN dep ensure --vendor-only -v
COPY server .
RUN go install .

FROM node:8-alpine as uibuilder
WORKDIR /usr/src/app
COPY ui/package.json ui/yarn.lock ./
RUN yarn install
COPY ui .
RUN yarn build

FROM alpine:3.8 as release
RUN apk add --no-cache ca-certificates \
    busybox-extras net-tools bind-tools
WORKDIR /frontend
COPY --from=serverbuilder /go/bin/server ./server
COPY --from=uibuilder /usr/src/app/build ./ui
# COPY ./server/templates ./templates
# COPY ./static ./static
EXPOSE 8080
ENTRYPOINT ["/frontend/server"]
