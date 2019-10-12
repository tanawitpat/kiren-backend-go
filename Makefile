install:
	glide update

build_local:
	docker-compose -f deployment/local/docker-compose.yml build

build_prd:
	docker build . -t kiren-api -f deployment/prd/Dockerfile

run: build_local
	docker-compose -f deployment/local/docker-compose.yml up -d
	docker logs -f kiren-api

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

unit_test:
	go test ./...

api_spec:
	go run swagger/main.go
