# kiren-backend-go

**Frontend Repository:** <https://github.com/tanawitpat/kiren-frontend>

## Project Maintainer

- Tanawit Pattanaveerangkoon <<tanawit.pat@gmail.com>>

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

Environment variables used in deployment will be imported from `.env`

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
# Environment variable: GCLOUD_PROJECT_ID
make deploy

# Deploy the service to AWS lambda
# Environment variables: SERVERLESS_STAGE, SERVERLESS_MYSQL_DB_NAME, SERVERLESS_MYSQL_HOST, SERVERLESS_MYSQL_PORT, SERVERLESS_MYSQL_USERNAME, SERVERLESS_MYSQL_PASSWORD
make deploy_lambda
```
