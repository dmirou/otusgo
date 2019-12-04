HAS_LINT := $(shell command -v golangci-lint;)

all: run

run: vendor lint
	@echo "+ $@"
	@go run ./cmd/main.go

vendor: bootstrap
	@echo "+ $@"
	@go mod tidy

bootstrap:
ifndef HAS_LINT
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
endif

lint: bootstrap
	@echo "+ $@"
	@golangci-lint run ./...

.PHONY: all \
		run \
		vendor \
		bootstrap \
		lint
