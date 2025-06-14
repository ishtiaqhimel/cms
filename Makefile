PROJECT_NAME := github.com/ishtiaqhimel/news-portal/cms
PKG_LIST := $(shell go list ${PROJECT_NAME}/... | grep -v /vendor/)

GO_PKG := github.com/ishtiaqhimel/news-portal/cms

.PHONY: help git-hook fmt vet lint swagger dep clean test

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z0-9-]+%?:.*#' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

git-hook:
	@git config --local core.hooksPath .githooks/

fmt: ## Run go fmt against code
	@go fmt ./...

vet: ## Run go vet against code
	@go vet ./...

GOLANGCI_LINT_VERSION := v2.1.6
LINTER_IMAGE ?= golangci/golangci-lint:$(GOLANGCI_LINT_VERSION)

lint: fmt vet ## Run go linter against code
	@echo "Running linter"
	@docker run           \
	-i                    \
	--rm                  \
	-v $$(pwd):/src       \
	-w /src               \
	$(LINTER_IMAGE)       \
	golangci-lint run --timeout=5m


########################
### DEVELOP and TEST ###
########################

dep%: git-hook ## Run dependent container(s) along with configuring consul based on the pattern matching passed with target
	# booting up dependency containers
	@docker-compose up -d consul

	# wait for consul container be ready
	@while ! curl --request GET -sL --url 'http://localhost:8500/' > /dev/null 2>&1; do printf .; sleep 1; done
	@echo


	@docker-compose up -d db_primary; \
	curl --request PUT --data-binary @config.local.yml http://localhost:8500/v1/kv/news-portal/cms;

dep: dep- ## Run default dependent container(s) along with configuring consul

serve%: dep% ## Run the project locally to develop (with hot reload!) based on the pattern matching passed with target
	@docker-compose up --build api

development-all: dep ## Run the project locally to develop (with hot reload!)
	@docker-compose build
	@docker-compose up

test: ## Run unittests
	@go test -coverprofile cov.out -short ${PKG_LIST}

coverage: ## Generate global code coverage report
	@go tool cover -func=cov.out

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)
	@docker-compose down
