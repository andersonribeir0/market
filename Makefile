clear:
	docker-compose -f docker-compose.yaml down

stop:
	docker-compose -f docker-compose.yaml stop

build-dev:
	docker-compose -f docker-compose.development.yaml build

build:
	docker-compose -f docker-compose.yaml build

dev:
	docker-compose -f docker-compose.development.yaml up -d
	docker exec -it market bash

dynamodb:
	docker-compose -f docker-compose.yaml up -d dynamodb
	curl localhost:8000 || sleep 5

app:
	docker-compose -f docker-compose.yaml up -d app
