FROM node:20.11.0

WORKDIR /usr/src/app

# グローバルパッケージのインストール先を変更
ENV NPM_CONFIG_PREFIX=/home/node/.npm-global

# 必要なディレクトリを作成し、権限を設定
RUN mkdir -p /home/node/.npm-global && \
    chown -R node:node /home/node/.npm-global && \
    chown -R node:node /usr/src/app

# nodeユーザーに切り替え
USER node

# PATHを更新
ENV PATH=/home/node/.npm-global/bin:$PATH

# 必要なパッケージをグローバルにインストール
RUN npm install -g next

# node_modulesディレクトリを作成
RUN mkdir -p /usr/src/app/node_modules

CMD ["npm", "run", "dev"]