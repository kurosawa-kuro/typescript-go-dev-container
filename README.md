# react-ts-go-docker
React TypeScript Go in Docker 

test
swagger
dev画面



make dev-backend


docker compose exec backend go test ./...

http://localhost:8000/swagger/index.html

http://localhost:3000

Pgadmin
http://localhost:5050/


Storybook
http://localhost:6006/

docker-compose down -v&& docker-compose up --build

	// SSH エージェントソケットをマウント
	"mounts": [
		"source=/run/host-services/ssh-auth.sock,target=/run/host-services/ssh-auth.sock,type=bind"
	],

	"remoteEnv": {
		"SSH_AUTH_SOCK": "/run/host-services/ssh-auth.sock"
	}


	docker compose -f .devcontainer/docker-compose.yml build --no-cache


# 1. すべてのコンテナを停止し削除（-vオプションでボリュームも削除）
docker-compose down -v

# 2. 未使用のコンテナ、ネットワーク、イメージ、ボリュームをすべて削除
docker system prune -a --volumes

# 3. （オプション）イメージを個別に削除する場合
docker rmi $(docker images -q)

# 4. （オプション）ボリュームを個別に削除する場合
docker volume rm $(docker volume ls -q)
================

cd .devcontainer/
docker-compose down -v


# Dockerの完全クリーンアップ
docker-compose down -v
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
docker system prune -a --volumes -f

# VSCode関連ファイルのクリーンアップ
rm -rf ~/.vscode-server
rm -rf ~/.vscode-remote-containers



cd .devcontainer/
docker-compose down -v

.devcontainerの構成をFrontendとBackendに合わせる
バックエンドの正しいテスト環境作成







cd .devcontainer/
# Dockerの完全クリーンアップ
# 1. 実行中のコンテナを停止
docker-compose down -v
docker stop $(docker ps -a -q)

# 2. 全てのコンテナを削除
docker rm $(docker ps -a -q)

# 3. 全てのイメージを削除
docker rmi $(docker images -q) -f

# 4. 全てのボリュームを削除
docker volume rm $(docker volume ls -q)

# 5. 全てのネットワークを削除
docker network prune -f

# 6. 未使用のシステムリソースを削除（キャッシュ含む）
docker system prune -a --volumes -f

# 7. Dev Container関連のVSCodeキャッシュを削除
rm -rf ~/.vscode-server
rm -rf ~/.vscode-remote-containers


# PostgreSQL基本設定
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=dev_db
POSTGRES_PORT=5432

# ホスト設定
# コンテナ間通信用（バックエンド→DB）
DOCKER_DATABASE_HOST=db              

# pgAdmin設定
PGADMIN_DEFAULT_EMAIL=admin@admin.com
PGADMIN_DEFAULT_PASSWORD=admin

# テストを実行
cd ./backend/src
go test ./handler/... ./test/... -v



docker exec -it test-postgres-db psql -U postgres -c "\l"
docker exec -it test-postgres-db psql -U postgres -c "CREATE DATABASE test_db;"

docker logs typescript-go-dev-container_devcontainer-db-1

host.docker.internal