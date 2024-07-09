# Stage 1: Build the Go application
FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o blog

# Stage 2: Run the Go application
FROM alpine:latest

WORKDIR /root/

# bashとその他の必要なツールをインストール
RUN apk --no-cache add bash curl && \
    curl -o wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh && \
    chmod +x wait-for-it.sh

COPY --from=builder /app/blog .

EXPOSE 8080

CMD ["bash", "./wait-for-it.sh", "mysql:3306", "--", "./blog"]
