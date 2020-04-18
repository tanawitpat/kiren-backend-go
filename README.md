# kiren-backend-go

## Project Maintainer

- Tanawit Pattanaveerangkoon <<tanawit.pat@gmail.com>>

## Project Overview

This project was developed for learning API development using Go.

`kiren-backend-go` has four endpoints (ping, getProduct, getProducts, getBestSellerProducts). It queries data from MySQL database and returns it in JSON format. You can view sample requests in `request.http`

Note: `Kiren` is the name of my family business. If you want a new Thai BBQ stove for your store, please let us know. (<https://www.facebook.com/kirenbbq>)

## Installation

### 1. Local host development

```bash
# 1. Install dependencies
make install

# 2. Start the kiren-mysql container (with initial data)
make run_mysql

# 3. Start the service (kiren-api)
go run main.go
```

### 2. Container native development

```bash
# 1. Install dependencies
make install

# 2. Run MySQL and api using Docker (the kiren-mysql and kiren-api containers will be started)
make run
```

## API Specification

The API specification is served by Swagger. It can be accessed via `make api_spec` command.

## Commands

### Google Kubernetes Engine Deployment

`GCLOUD_PROJECT_ID` must be defined in the `.env` file. It will be used for GKE deployment.

### Serverless Deployment

The following environment variables must be defined in the `.env` file. It will be used for serverless configuration file generation.

- SERVERLESS_STAGE
- SERVERLESS_MYSQL_DB_NAME
- SERVERLESS_MYSQL_HOST
- SERVERLESS_MYSQL_PORT
- SERVERLESS_MYSQL_USERNAME
- SERVERLESS_MYSQL_PASSWORD

```bash
# Install libraries
make install

# Build docker image
make build

# Start the kiren-mysql container
make run_mysql

# Start all mandatory containers (kiren-mysql and kiren-api)
make run

# Run unit tests
make unit_test

# Preview API specification
make api_spec

# Deploy the service to GKE
make deploy

# Deploy the service to AWS lambda
make deploy_lambda
```
