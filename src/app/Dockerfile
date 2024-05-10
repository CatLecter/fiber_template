FROM golang:1.22-bookworm as builder

RUN mkdir "/tmp/build"

WORKDIR /tmp/build

COPY . .

RUN go mod tidy &&  \
    go mod download &&  \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main .

FROM debian:12-slim

WORKDIR /usr/src/app

COPY docs .

COPY --from=builder /tmp/build .

CMD /usr/src/app/main
