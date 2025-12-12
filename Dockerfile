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

RUN echo "=== Checking dist structure ===" && \
    ls -la dist/ && \
    echo "=== Checking dist/assets ===" && \
    ls -la dist/assets/ || echo "No assets in dist"

FROM alpine:latest

WORKDIR /app

COPY --from=backend-builder /server /app/server
COPY --from=frontend-builder /app/testops-copilot/dist /app/static

RUN echo "=== Final container structure ===" && \
    ls -la /app/static/ && \
    echo "=== Static assets ===" && \
    ls -la /app/static/assets/ || echo "No assets folder"

RUN chmod +x /app/server

CMD ["/app/server"]