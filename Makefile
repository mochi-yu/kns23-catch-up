COMPOSE := docker-compose -f docker-compose.local.yaml --env-file .env.local
include .env.local

.PHONY: build up down migrate console db seed init setup
setup:
	pnpm install -g dynamodb-admin
build: 
	$(COMPOSE) build
up: # サーバーを起動する
	DYNAMO_ENDPOINT=http://localhost:${DYNAMO_DB_PORT} dynamodb-admin -p ${DYNAMO_DB_GUI_PORT} & \
	$(COMPOSE) up
up-d: # サーバーをバックグラウンドで起動する
	$(COMPOSE) up -d
down: # サーバーを落とす
	$(COMPOSE) down
migrate: # 差分マイグレートを実行
	# TODO implement
	# $(COMPOSE) exec app rails db:migrate
console: # Railsコンソールを開く
	# TODO implement
	# $(COMPOSE) exec app rails console
db: # DBコンソールを開く
	# TODO implement
	# $(COMPOSE) exec app rails db
seed: # DBにseed適用する（アプリケーションサーバーのみ落とし、DB再構築）
	# TODO implement
	# docker compose up -d
	# docker-compose rm -fsv app
	# docker compose exec rails db:migrate:reset
	# docker compose exec rails db:seed
	# docker compose up -d
init: # DBを空にする（アプリケーションサーバーのみ落とし、DB初期化）
	# TODO implement
	# docker compose up -d
	# docker-compose rm -fsv app
	# docker compose exec rails db:migrate:reset
	# docker compose up -d