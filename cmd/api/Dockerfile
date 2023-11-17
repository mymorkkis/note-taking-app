FROM golang:1.21.3-alpine3.18

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-build.cmd", "go build -o ./tmp/main ./cmd/api"]
