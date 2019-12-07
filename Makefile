#!make
include .env
export

# Install package dependencies
install:
	glide update

# Environment setup
build_local:
	docker-compose -f deployment/local/docker-compose.yml build

build_prd:
	docker build . -t kiren-api -f deployment/prd/Dockerfile

run: build_local
	docker-compose -f deployment/local/docker-compose.yml up -d
	docker logs -f kiren-api

# Kubernetes deployment
release: build_prd
	docker tag kiren-api asia.gcr.io/${GCLOUD_PROJECT_ID}/kiren-api
	docker push asia.gcr.io/${GCLOUD_PROJECT_ID}/kiren-api

connect_gke:
	gcloud container clusters get-credentials kiren --zone asia-southeast1-a --project ${GCLOUD_PROJECT_ID}

generate_deployment:
	cp deployment/prd/api-deployment.yaml.template deployment/prd/api-deployment.yaml
	sed -i "" 's/$${GCLOUD_PROJECT_ID}/'${GCLOUD_PROJECT_ID}'/g' deployment/prd/api-deployment.yaml

deploy: generate_deployment
	kubectl apply -f deployment/prd/api-deployment.yaml

apply_secret:
	kubectl apply -f deployment/prd/api-secrets.yaml

# AWS Lambda
clear_bin: 
	rm -rf ./bin

generate_serverless_deployment:
	cp serverless.yml.template serverless.yml
	sed -i "" 's/$${MYSQL_DB_NAME}/'${MYSQL_DB_NAME}'/g' serverless.yml
	sed -i "" 's/$${MYSQL_HOST}/'${MYSQL_HOST}'/g' serverless.yml
	sed -i "" 's/$${MYSQL_PORT}/'${MYSQL_PORT}'/g' serverless.yml
	sed -i "" 's/$${MYSQL_USERNAME}/'${MYSQL_USERNAME}'/g' serverless.yml
	sed -i "" 's/$${MYSQL_PASSWORD}/'${MYSQL_PASSWORD}'/g' serverless.yml

build_lambda:
	GOOS=linux go build -o bin/ping lambda/ping/main.go
	GOOS=linux go build -o bin/getProduct lambda/getProduct/main.go
	GOOS=linux go build -o bin/getProducts lambda/getProducts/main.go

pack_lambda:
	zip bin/ping.zip bin/ping
	zip bin/getProduct.zip bin/getProduct
	zip bin/getProducts.zip bin/getProducts

deploy_lambda: generate_serverless_deployment clear_bin build_lambda
	sls deploy --verbose

# Development utils
unit_test:
	go test ./...

api_spec:
	go run swagger/main.go
