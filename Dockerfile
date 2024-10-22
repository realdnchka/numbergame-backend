FROM golang:1.23.2-bookworm as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

ENTRYPOINT [ "go run main.go" ]
