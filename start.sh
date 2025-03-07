#!/bin/sh

export APP_HOME=/app/duck

# 启动 duck-server
cd ${APP_HOME}/server
${APP_HOME}/server/duck &

# 启动 duck-site
cd ${APP_HOME}/site
node ${APP_HOME}/site/.output/server/index.mjs &

# 保持容器运行
wait
