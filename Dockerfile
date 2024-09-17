FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o minesweeper ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/minesweeper .

EXPOSE 3000

CMD ["./minesweeper"]