# Stage 1: Build the Go application
FROM golang:1.18-alpine AS builder

WORKDIR /app

# 必要なツールをインストール
RUN apk --no-cache add git curl

# 依存関係をインストール
COPY go.mod go.sum ./
RUN go mod download

# airを特定バージョンでインストール
RUN go install github.com/cosmtrek/air@v1.29.0

# ソースコードをコピー
COPY . .

# Stage 2: Run the Go application with Air
FROM golang:1.18-alpine

WORKDIR /app

# 必要なツールをインストール
RUN apk --no-cache add bash curl

# wait-for-it.sh スクリプトをダウンロード
RUN curl -o /usr/local/bin/wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh && \
    chmod +x /usr/local/bin/wait-for-it.sh

# ビルドされたアプリケーションと必要なツールをコピー
COPY --from=builder /app /app
COPY --from=builder /go/bin/air /usr/local/bin/air

EXPOSE 8080

CMD ["bash", "wait-for-it.sh", "mysql:3306", "--", "air"]
