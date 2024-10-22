FROM golang:1.23.2-alpine as builder

WORKDIR /app

COPY go.* .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

FROM alpine:latest
COPY --from=builder /app/server /app/server
EXPOSE 80

CMD [ "/app/server" ]
