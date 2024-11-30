# システムのアップデート
sudo yum update -y

# Gitのインストール
sudo yum install -y git
git --version

# Dockerのインストール
sudo yum install -y docker
sudo systemctl start docker
sudo systemctl enable docker
sudo usermod -aG docker ec2-user

# Docker Composeのインストール
# バージョン2.xの場合
sudo mkdir -p /usr/local/lib/docker/cli-plugins/
sudo curl -SL https://github.com/docker/compose/releases/latest/download/docker-compose-linux-x86_64 -o /usr/local/lib/docker/cli-plugins/docker-compose
sudo chmod +x /usr/local/lib/docker/cli-plugins/docker-compose

# インストールの確認
docker --version
docker compose version

# 注意: 変更を反映するためにシステムの再起動かログアウト/ログインが必要
# sudo reboot
# または
exit  # 再度SSHで接続