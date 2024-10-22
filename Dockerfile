FROM golang:1.23.2-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./
RUN ls

CMD [ "go run main.go" ]
