FROM golang:1.24.4-alpine3.22 as builder

RUN mkdir "/tmp/build"

WORKDIR /tmp/build

COPY ../docs docs
COPY ../cmd cmd
COPY ../configs configs
COPY ../internal internal
COPY ../vendor vendor
COPY ../go.sum go.sum
COPY ../go.mod go.mod

RUN go mod tidy
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=vendor -o main cmd/main.go

FROM alpine:3.22

WORKDIR /usr/src/app

COPY ../docs docs

COPY --from=builder /tmp/build .

CMD /usr/src/app/main
