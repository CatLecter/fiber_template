FROM golang:1.22-bookworm as builder

RUN mkdir "/tmp/build"

WORKDIR /tmp/build

COPY ../api api
COPY ../app app
COPY ../cmd cmd
COPY ../configs configs
COPY ../internal internal
COPY ../vendor vendor
COPY ../go.sum go.sum
COPY ../go.mod go.mod

RUN go mod tidy
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=vendor -o main cmd/main.go

FROM debian:12-slim

WORKDIR /usr/src/app

COPY ../api api

COPY --from=builder /tmp/build .

CMD /usr/src/app/main
