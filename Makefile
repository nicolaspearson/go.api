.PHONY: help

help: ## this help output
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

db-start: ## start the database
	docker-compose up -d postgres

docker-build: ## build the docker image
	docker-compose build --force-rm api

docker-run: ## run the docker image
	docker-compose up

go-build: ## build the go app
	go build -o tmp/api ./cmd/api/main.go

go-install: ## install the dependencies
	go install ./cmd/api/main.go

go-run: ## run the go app
	go run ./cmd/api/main.go

go-test: ## execute tests
	go test ./cmd/api/... -cover

swag-init: ## generate swagger docs
	swag init --dir ./cmd/api --output ./cmd/api/docs --parseDependency --parseInternal
