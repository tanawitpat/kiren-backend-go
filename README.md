# kiren-backend-go

**Frontend Repository:** <https://github.com/tanawitpat/kiren-frontend>

## Project Maintainer

- Tanawit Pattanaveerangkoon <<tanawit.pat@gmail.com>>

## API Specification

The API specification is served by Swagger. It can be accessed via `make api_spec` command.

## Commands

The environment variables will be imported from '.env'.

```bash
# Install libraries
make install

# Build docker image
make build

# Run the service
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
