# go.api

A **golang** RESTful API starter project.

The starter project includes the following out of the box:

- Authorization
- Error handling
- Input validation
- Swagger documentation
- Database integration (go-pg + PostgreSQL)
- Containerization via Docker
- Unit tests
- Integration tests
- Complete CI workflow (build -> test -> publish)

## Getting Started

Run the application:
```bash
make go-run
```

Execute the tests:
```bash
make go-test
```

### Docker

Build the Docker image:
```bash
make docker-build
```

Run the Docker image:
```bash
make docker-run
```

## Directory Structure

- `build` contains all shell scripts needed to build and test the application.
- `cmd` contains the source code! By convention, the source directory is named `cmd`, within, there is another directory with the name of the project - in this case `api`. This directory contains the `main.go` file that runs the *Go* application. The rest of the source is further divided into modules in this directory.
- `config` contains files with all required environment variables.
- `pkg` contains a *Go* package that only contains the global app version string. This is substituted for the actual version computed from the commit hash during build.

## Contribution Guidelines

Never commit directly to master, create a feature branch and submit a pull request.
