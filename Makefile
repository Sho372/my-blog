# 定義
DOCKER_COMPOSE = docker-compose

# コンテナの再起動
.PHONY: restart
restart:
	$(DOCKER_COMPOSE) restart

# DBコンテナの初期化
.PHONY: db-init
db-init:
	$(DOCKER_COMPOSE) down -v
	$(DOCKER_COMPOSE) up -d mysql

# appコンテナのビルドと再起動
.PHONY: app-rebuild
app-rebuild:
	$(DOCKER_COMPOSE) build app
	$(DOCKER_COMPOSE) up -d app

# 全コンテナのビルドと再起動
.PHONY: rebuild
rebuild:
	$(DOCKER_COMPOSE) down
	$(DOCKER_COMPOSE) up --build -d

# ログの表示
.PHONY: logs
logs:
	$(DOCKER_COMPOSE) logs -f

# 停止
.PHONY: stop
stop:
	$(DOCKER_COMPOSE) down

# クリーンアップ
.PHONY: clean
clean:
	$(DOCKER_COMPOSE) down -v
