FROM golang:1.20.3-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY ./src ./src

RUN go build -o ./bin/go-doit ./src/main.go