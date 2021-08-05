clear:
	find . -name "*.log" -type f -delete
	docker-compose -f docker-compose.yaml down
	docker-compose -f docker-compose.test.yaml down

stop:
	docker-compose -f docker-compose.yaml stop
	docker-compose -f docker-compose.test.yaml stop

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
	docker-compose -f docker-compose.test.yaml run app go test -covermode=count -coverprofile=coverage.out -json > test_report.json ./...
	docker-compose -f docker-compose.test.yaml run app go tool cover -html=coverage.out -o coverage.html

app:
	docker-compose -f docker-compose.yaml up -d app

sonar:
	docker run -it \
	--rm \
	--net market_default \
	-e SONAR_HOST_URL="http://sonarqube:9000" \
	-e SONAR_LOGIN="b100df0f1d38ad87a12951b16f7e17dc5309514e" \
	-v $(shell pwd):/usr/src  \
	sonarsource/sonar-scanner-cli \
	-D sonar.projectKey=com.github.andersonribeir0.market \
    -D sonar.projectName=market \
	-D sonar.projectBaseDir=. \
    -D sonar.projectVersion=1.0 \
    -D sonar.sourceEncoding=UTF-8 \
	-D sonar.sources=. \
	-D sonar.exclusions=**/*_test.go \
	-D sonar.tests=. \
	-D sonar.test.inclusions=**/*_test.go \
	-D sonar.go.coverage.reportPaths=coverage.out \
	-D sonar.go.test.reportPaths=test_report.json
	

