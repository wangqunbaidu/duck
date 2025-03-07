# server builder

FROM golang:1.24 AS server_builder

ENV APP_HOME=/code/duck/server
WORKDIR "$APP_HOME"

COPY ./server ./
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN CGO_ENABLED=0 go build -v -o duck main.go && chmod +x duck


# site builder
FROM node:20-alpine AS site_builder

ENV APP_HOME=/code/duck/site
WORKDIR "$APP_HOME"

COPY ./site ./
# RUN npm install -g pnpm --registry=https://registry.npmmirror.com
# RUN pnpm install --registry=https://registry.npmmirror.com
RUN npm install -g pnpm
RUN pnpm install
RUN pnpm build:docker


# admin builder
FROM node:20-alpine AS admin_builder

ENV APP_HOME=/code/duck/admin
WORKDIR "$APP_HOME"

COPY ./admin ./
# RUN npm install -g pnpm --registry=https://registry.npmmirror.com
# RUN pnpm install --registry=https://registry.npmmirror.com
RUN npm install -g pnpm
RUN pnpm install
RUN pnpm build:docker

# run
FROM node:20-alpine

ENV APP_HOME=/app/duck
WORKDIR "$APP_HOME"

COPY --from=server_builder /code/duck/server/duck ./server/duck
COPY --from=server_builder /code/duck/server/*.yaml ./server/
COPY --from=server_builder /code/duck/server/*.yml ./server/
COPY --from=server_builder /code/duck/server/migrations ./server/migrations
COPY --from=site_builder /code/duck/site/.output ./site/.output
COPY --from=site_builder /code/duck/site/node_modules ./site/node_modules
COPY --from=admin_builder /code/duck/admin/dist ./server/admin

COPY ./start.sh ${APP_HOME}/start.sh
RUN chmod +x ${APP_HOME}/start.sh

EXPOSE 8082 3000

CMD ["./start.sh"]
