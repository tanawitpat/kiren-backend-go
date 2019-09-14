install:
	glide update

build:
	docker-compose -f deployment/local/docker-compose.yml build

run: build
	docker-compose -f deployment/local/docker-compose.yml up -d
	docker logs -f kiren-api

unit_test:
	go test ./...