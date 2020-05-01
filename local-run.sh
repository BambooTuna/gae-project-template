#!/bin/bash

cd `dirname $0`

# ミドルウェアを取得
if cd middleware; then
  git pull;
  cd ../
else
  git clone https://github.com/BambooTuna/middleware.git middleware
fi

if type "docker-compose" > /dev/null 2>&1; then
  docker-compose -f docker-compose-local.yaml up --build -d
else
  docker run \
  --rm -v /var/run/docker.sock:/var/run/docker.sock \
  -v "$PWD:/$PWD" -w="/$PWD" \
  docker/compose:1.22.0 \
  -f docker-compose-local.yaml up --build -d
fi
