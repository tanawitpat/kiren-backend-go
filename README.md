# kiren-backend-go

**Frontend URL:** <https://kiren.tanawitp.me>

## Project Maintainer

- Tanawit Pattanaveerangkoon <<tanawit.pat@gmail.com>>

## API Specification

The API specification is served by Swagger. It can be accessed via `make api_spec` command.

## Commands

```bash
# Install libraries
make install

# Build docker image
make build

# Run the service
make run

# Run unit tests
make unit_test

# Deploy
make deploy GCLOUD_PROJECT_ID=<GCLOUD_PROJECT_ID>

# Preview API specification
make api_spec
```
