# How To Run

## Pre-requisites

- Platform:
  - Docker & Docker Compose: version 27.3.1
  - Python: version 3.13.2
  - Go: version 1.24.1
- Tools:
  - [pre-commit](https://pre-commit.com/): version 4.0.1
  - [golangci-lint](https://golangci-lint.run/): version 2.0.2
  - [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports): version 0.31.0
  - [go-testfixtures] (https://github.com/go-testfixtures/testfixtures): version 3.14.0

You can install these tools with:

```bash
make install-prerequisites
```

To install the pre-commit hooks, run:

```bash
make install-precommit-hooks
```

## Deploy Infrastructure

To deploy the infrastructure, run:

```bash
make up
```

This command will start the Docker containers defined in the `docker-compose.yml` file, setting up the necessary services including:

- PostgreSQL database: stores the tenant's control plane data, such as tenant's connector settings, alert configurations, etc.
- Kafka + Kafka Connect + Schema Registry: used to stream data from and to tenant's datasources.
- Kafbat UI: a web-based user interface for managing Kafka topics and connectors, useful for debugging and monitoring.

## Load Mock Data

To load mock data into the PostgreSQL database, run:

```bash
make load-mock-data
```

This command will use testfixtures to load mock data defined in `testdata/fixtures` directory into the PostgreSQL database.

## Run the API

You need to create environment file `.env.local` in the root directory of the project. You can use `.env.example` as a template:

```bash
cp .env.example .env.local
```

Should you didn't change the default credentials in the `docker-compose.yml` file, you don't need to change anything in the `.env.local` file.

To run the API locally:

```bash
make run-local
```

## Summary

```bash
make install-prerequisites
make install-precommit-hooks
make up
make load-mock-data
make run-local
```
