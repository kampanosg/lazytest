.DEFAULT_GOAL := build

.PHONY: build
build: ## Build the binary
	go build -o lazytest ./main.go

.PHONY: run
run: build ## Run the binary
	./lazytest

.PHONY: test
test: ## Run the tests
	go test ./...

.PHONY: mockgen
mockgen: ## Generate the mocks
	go generate ./...

.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
