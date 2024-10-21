FROM golang:1.23.2-alpine3.20

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-build.cmd", "go build -o ./tmp/main ./cmd/api"]
