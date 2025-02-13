FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o todo-app ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/todo-app .

COPY internal/config /root/internal/config

COPY .env .

CMD ["./todo-app"]
