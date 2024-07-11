# Stage 1: Base dependencies
FROM golang:1.22-alpine AS base

# 必要なツールをインストール
RUN apk --no-cache add git curl gcc libc-dev

# 作業ディレクトリを設定
WORKDIR /app

# Goモジュールの依存関係をコピーしてダウンロード
COPY go.mod go.sum ./
RUN go mod download

# airを特定バージョンでインストール
RUN go install github.com/air-verse/air@v1.52.3

# Stage 2: Build the Go application
FROM base AS builder

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN CGO_ENABLED=1 go build -o blog .

# Stage 3: Final image
FROM golang:1.22-alpine

# 必要なツールをインストール
RUN apk --no-cache add bash curl gcc libc-dev

# 作業ディレクトリを設定
WORKDIR /app

# wait-for-it.sh スクリプトをダウンロード
RUN curl -o /usr/local/bin/wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh && \
    chmod +x /usr/local/bin/wait-for-it.sh

# ビルドされたアプリケーションと必要なツールをコピー
COPY --from=builder /app /app
COPY --from=builder /go/bin/air /usr/local/bin/air

# 環境変数を設定
ENV CGO_ENABLED=1

EXPOSE 8080

CMD ["bash", "wait-for-it.sh", "mysql:3306", "--", "air"]
