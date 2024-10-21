FROM golang:1.23.2-bookworm as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o server

FROM debian:bookworm-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /app/server
EXPOSE 80

CMD ["/app/server"]