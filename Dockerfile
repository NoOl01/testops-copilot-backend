FROM golang:1.24-alpine AS backend-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /server ./cmd/app

FROM node:24-alpine AS frontend-builder

WORKDIR /app

RUN apk add --no-cache git

RUN git clone https://github.com/K1butsujiMuzan/testops-copilot.git

WORKDIR /app/testops-copilot

RUN npm install
RUN npm run build

FROM alpine:latest

WORKDIR /app

COPY --from=backend-builder /server /app/server
COPY --from=frontend-builder /app/testops-copilot/dist /app/static

RUN chmod +x /app/server

CMD ["/app/server"]
