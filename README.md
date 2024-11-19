# react-ts-go-docker
React TypeScript Go in Docker 

docker-compose down


docker-compose down -v
docker-compose up -d --build backend
docker logs -f backend


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

# Dockerの完全クリーンアップ
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
docker system prune -a --volumes -f

# VSCode関連ファイルのクリーンアップ
rm -rf ~/.vscode-server
rm -rf ~/.vscode-remote-containers