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
  - [goose] (https://github.com/pressly/goose): version 3.24.2

You can install these tools with:

```bash
make install-prerequisites
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
make load-mock-db
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
make up
make load-mock-db
make run-local
```

# Structure Explanation

## Overall Layout

The project follows a standard Go project structure with additional directories tailored for specific functionalities. It employs a layered architecture, separating concerns into domain, application, and infrastructure layers. Key directories include:

- **`.` (root)**: Houses configuration, documentation, and build-related files.
- **`cmd`**: Contains the application's entry points.
- **`internal`**: Encapsulates private application code, organized by functionality (e.g., controllers, services, repositories).
- **`migrations`**: Manages database schema changes.
- **`scripts`**: Holds utility scripts.
- **`testdata`**: Stores test fixtures.

## Directory Descriptions

### `.`

- **`LICENSE`**: Contains the project's license information.
- **`README.md`**: Provides project documentation and an overview.
- **`go.mod`** and **`go.sum`**: Manage project dependencies using Go modules.
- **`makefile`**: Includes utility commands for building, testing, and running the application.
- **`docker-compose.yaml`**: Defines services, networks, and volumes for Docker containers, facilitating local development and testing.
- **`.env.example`**: A template for environment variables, providing a reference for configuration.

### `cmd`

- **`api/`**: Contains the main entry point for the API server.
  - **`main.go`**: The application's entry point, where execution begins.
  - **`app.go`**: Initializes the application, setting up services, middleware, and dependencies.
  - **`config.go`**: Manages configuration loading and setup, such as environment variables or config files.

### `internal`

This directory holds private application code, preventing external imports, and is organized into subdirectories by functionality.

- **`client/`**: Implements clients for interacting with external services.
- **`controller/`**: Manages incoming HTTP requests and routes them to services.
- **`db/`**: Sets up database connections.
- **`domain/`**: Defines core business entities and models.
- **`dto/`**: Data Transfer Objects.
- **`enum/`**: Defines enumerations or constants.
- **`logging/`**: Configures and provides logging utilities.
- **`mapper/`**: Contains functions to map between data structures (e.g., domain models to DTOs and vice versa).
- **`middleware/`**: Implements HTTP middleware functions.
- **`repository/`**: Encapsulates data access logic for interacting with the database.
- **`service/`**: Houses business logic, orchestrating operations between controllers and repositories.

### `migrations`

- **`strimq.sql`**: SQL scripts for database schema migrations.

### `scripts`

- Contains random ad-hoc scripts for various tasks.

### `testdata`

- **`fixtures/`**: YAML files for testing or seeding data.

## Naming Conventions

- **Package Names**: singular, short and reflective of their purpose (e.g., `controller`, `service`, `repository`).
- **Directories**: Typically named the same as the package they represent (e.g., `controller/`, `service/`, `repository/`).
- **Files**: `<functionality>_<package>.go` format. A few notes:
  - A file may contain multiple related functions or types, but the file name should reflect the primary functionality or type it provides.
  - Domain model files don't need prepending with `_<package>` (e.g., `domain/user.go`, `domain/tenant.go`).

## Architecture

I'm following a half-baked clean architecture (no over-evangelization), with clear separation of concerns:

- **DTO**: Data Transfer Objects for request and response payloads (`dto/`).
- **Presentation**: HTTP controllers and middleware (`controller/`, `middleware/`).
- **Domain**: Core business logic and models (`domain/`).
- **Application**: Business rules and services (`service/`).
- **Infrastructure**: External interactions like databases and clients (`repository/`, `client/`).
