FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /server ./cmd/app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /server /app/server

RUN chmod +x server

CMD ["/app/server"]