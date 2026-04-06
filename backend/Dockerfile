FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
RUN go install -tags "postgres" github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.0
COPY . .
RUN go build -o backend ./cmd/server

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/backend .
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY migrations ./migrations
ENV PATH=/usr/local/bin:$PATH

EXPOSE 8080

# Set default command: run migrations then start backend
CMD sh -c "echo \$DATABASE_URL && ls -l ./migrations && migrate -path ./migrations -database \$DATABASE_URL up && ./backend"