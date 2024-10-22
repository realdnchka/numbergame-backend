FROM golang:1.23.2-bookworm as builder

RUN go mod download

ENTRYPOINT [ "go run main.go" ]
