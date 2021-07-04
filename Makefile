clear:
	find . -name "*.log" -type f -delete
	docker-compose -f docker-compose.yaml down

stop:
	docker-compose -f docker-compose.yaml stop

run:
	docker-compose -f docker-compose.yaml up -d

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

createTable:
	docker-compose -f docker-compose.yaml run app go run main.go createTable

deleteTable:
	docker-compose -f docker-compose.yaml run app go run main.go deleteTable

importCsv:
	docker-compose -f docker-compose.yaml run app go run main.go importCsv

test:
	docker-compose -f docker-compose.yaml run app go test -covermode=count -coverprofile=coverage.out ./...
	docker-compose -f docker-compose.yaml run app go tool cover -html=coverage.out -o coverage.html

app:
	docker-compose -f docker-compose.yaml up -d app
