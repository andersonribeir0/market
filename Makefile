init:
	docker volume create --name=dynamodb_data
	docker volume create --name=market_data

clear:
	docker-compose -f docker-compose.yaml down

stop:
	docker-compose -f docker-compose.yaml stop

build:
	docker-compose -f docker-compose.development.yaml build

exec-dev:
	docker-compose -f docker-compose.development.yaml up -d
	docker exec -it market bash

prepare:
	docker-compose -f docker-compose.yaml up -d dynamodb
	curl localhost:8000 || sleep 15

