# main project name for naming main build
PROJECT_NAME=growth-place

.PHONY: tools
tools: ## Install all needed tools, e.g. for static checks
	@echo Installing tools from tools.go
	@cd tools; grep '_ "' tools.go | grep -o '"[^"]*"' | xargs -tI % go install %

.PHONY: build
build: ## build project to binary file
	go build -o ${PWD}/${PROJECT_NAME} ${PWD}/cmd

.PHONY: run
run: ## run project binary file
	./${PROJECT_NAME}

.PHONY: dependencies
dependencies: ## manage project dependencies
	go mod tidy

.PHONY: download-dependencies
download-dependencies: ## download all project dependencies from go.mod
	go mod download

.PHONY: migrate-up
migrate-up: ## up all migrates on db
	sql-migrate up

.PHONY: migrate-down
migrate-down: ## back on one migrate step on db
	sql-migrate down

.PHONY: docs
docs: ## generate swagger documentation
	swag init -g cmd/main.go

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[95m%-30s\033[0m %s\n", $$1, $$2}'
