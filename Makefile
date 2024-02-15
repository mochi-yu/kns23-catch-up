COMPOSE := docker-compose -f docker-compose.local.yaml --env-file .env.local
include .env.local

.PHONY: build up down migrate console db seed init setup
setup:
	pnpm install -g dynamodb-admin
build: 
	$(COMPOSE) build
up: # サーバーを起動する
	DYNAMO_ENDPOINT=http://localhost:$(DYNAMO_DB_PORT) dynamodb-admin -p $(DYNAMO_DB_GUI_PORT) & \
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


# DynamoDB Local用のコマンド
.PHONY: create_table_local drop_table_local
create_table_local:
	aws-vault exec $(AWS_ACCOUNT_NAME) -- aws dynamodb \
		create-table --table-name users \
		--attribute-definitions \
			AttributeName=firebase_id,AttributeType=S \
		--key-schema \
			AttributeName=firebase_id,KeyType=HASH \
		--billing-mode=PAY_PER_REQUEST \
		--endpoint-url http://localhost:$(DYNAMO_DB_PORT)
	aws-vault exec $(AWS_ACCOUNT_NAME) -- aws dynamodb \
		create-table --table-name posts \
		--attribute-definitions \
			AttributeName=post_id,AttributeType=S \
		--key-schema \
			AttributeName=post_id,KeyType=HASH \
		--billing-mode=PAY_PER_REQUEST \
		--endpoint-url http://localhost:$(DYNAMO_DB_PORT)
	aws-vault exec $(AWS_ACCOUNT_NAME) -- aws dynamodb \
		create-table --table-name temp_users \
		--attribute-definitions \
			AttributeName=firebase_id,AttributeType=S \
		--key-schema \
			AttributeName=firebase_id,KeyType=HASH \
		--billing-mode=PAY_PER_REQUEST \
		--endpoint-url http://localhost:$(DYNAMO_DB_PORT)

drop_table_local:
	aws-vault exec $(AWS_ACCOUNT_NAME) -- aws dynamodb \
		delete-table --table-name users \
		--endpoint-url http://localhost:$(DYNAMO_DB_PORT)
	aws-vault exec $(AWS_ACCOUNT_NAME) -- aws dynamodb \
		delete-table --table-name posts \
		--endpoint-url http://localhost:$(DYNAMO_DB_PORT)
	aws-vault exec $(AWS_ACCOUNT_NAME) -- aws dynamodb \
		delete-table --table-name temp_users \
		--endpoint-url http://localhost:$(DYNAMO_DB_PORT)
