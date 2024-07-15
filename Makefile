# 定義
DOCKER_COMPOSE = docker compose
APP_CONTAINER = blog_app

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

# 全テストの実行
.PHONY: test
test:
	docker exec $(APP_CONTAINER) sh -c "go test ./..."

# 特定のテストファイルを実行
.PHONY: test-file
test-file:
	@test -n "$(FILE)" || (echo "FILE is not set"; exit 1)
	docker exec $(APP_CONTAINER) sh -c "go test $(FILE)"

# 特定のテスト関数を実行
.PHONY: test-func
test-func:
	@test -n "$(FUNC)" || (echo "FUNC is not set"; exit 1)
	docker exec $(APP_CONTAINER) sh -c "go test -run $(FUNC)"
