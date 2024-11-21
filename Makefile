.PHONY: up down rebuild logs ps clean help test-frontend test-frontend-watch test-backend exec-frontend exec-backend exec-db

# デフォルトのターゲット
.DEFAULT_GOAL := help

# 変数定義
DC := docker-compose

###################
# 基本コマンド
###################
up: ## コンテナを起動
	$(DC) up -d

down: ## コンテナを停止
	$(DC) down

ps: ## 実行中のコンテナを表示
	$(DC) ps

logs: ## コンテナのログを表示
	$(DC) logs -f

clean: ## コンテナ、ボリューム、ネットワークを削除
	$(DC) down -v --rmi all --remove-orphans

###################
# クリーンアップ関連
###################
clean: ## コンテナ、ボリューム、ネットワークを削除
	$(DC) down -v --rmi all --remove-orphans

clean-all: clean ## Dockerシステム全体のクリーンアップ（すべてのコンテナ、イメージ、ボリュームを削除）
	docker system prune -a --volumes -f

clean-deep: ## 完全クリーンアップ（Dev Container環境を完全にリセット）
	cd .devcontainer && \
	$(DC) down -v && \
	docker stop $$(docker ps -a -q) 2>/dev/null || true && \
	docker rm $$(docker ps -a -q) 2>/dev/null || true && \
	docker rmi $$(docker images -q) -f 2>/dev/null || true && \
	docker volume rm $$(docker volume ls -q) 2>/dev/null || true && \
	docker network prune -f && \
	docker system prune -a --volumes -f && \
	rm -rf ~/.vscode-server && \
	rm -rf ~/.vscode-remote-containers

reset: clean-deep ## 完全なリセット（すべてを削除して再ビルド）
	$(DC) up --build -d

###################
# ビルド関連
###################
rebuild: down ## コンテナを再ビルドして起動
	$(DC) up --build

rebuild-frontend: down ## フロントエンドのコンテナを再ビルドして起動
	$(DC) up --build frontend

rebuild-backend: down ## バックエンドのコンテナを再ビルドして起動
	$(DC) up --build backend

###################
# コンテナ接続
###################
exec-frontend: ## フロントエンドのコンテナに入る
	$(DC) exec frontend sh

exec-backend: ## バックエンドのコンテナに入る
	$(DC) exec backend sh

exec-db: ## DBのコンテナに入る
	$(DC) exec db sh

###################
# テスト関連
###################
test-frontend: ## フロントエンドのテストを実行
	$(DC) exec frontend npm test

test-frontend-watch: ## フロントエンドのテストをウォッチモードで実行
	$(DC) exec frontend npm run test:watch

test-backend: ## バックエンドのテストを実行
	cd backend/src && go test ./handler/... ./test/... -v

test-backend-watch: ## バックエンドのテストをウォッチモードで実行
	cd backend/src && go test ./handler/... ./test/... -v -watch

test: test-frontend test-backend

# 開発用コマンド
###################
dev-frontend: ## フロントエンドの開発サーバーを起動
	cd frontend && npm run dev

dev-backend: ## バックエンドの開発サーバーを起動（ホットリロード）
	cd backend/src && air -c ../.air.toml

dev: dev-backend dev-frontend 

###################
# ヘルプ
###################
help: ## このヘルプを表示
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'