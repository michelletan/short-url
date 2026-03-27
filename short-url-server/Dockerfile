# Step 1: Build Stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o short-url-server ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/short-url-server .

EXPOSE 8080

CMD ["./short-url-server"]