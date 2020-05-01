#!/bin/bash

# 第一引数はgitリポジトリを指定(以下例)
# sh deploy.sh https://github.com/dockersamples/example-voting-app

#rm -rf app

# 初回のみClone、以降はPullする
if cd app; then
  git pull;
else
  git clone $1 app
  cd app
fi

# ミドルウェアを取得
if cd middleware; then
  git pull;
  cd ../
else
  git clone https://github.com/BambooTuna/middleware.git middleware
fi

# 全て削除
#docker run \
#--rm -v /var/run/docker.sock:/var/run/docker.sock \
#-v "$PWD:/$PWD" -w="/$PWD" \
#docker/compose:1.22.0 \
#down --rmi all --volumes

docker run \
--rm -v /var/run/docker.sock:/var/run/docker.sock \
-v "$PWD:/$PWD" -w="/$PWD" \
docker/compose:1.22.0 \
up --build -d
