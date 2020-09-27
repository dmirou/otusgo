PROJECT = github.com/dmirou/otusgo/template
HAS_LINT := $(shell command -v golangci-lint;)

all: vendor fmt test lint

vendor: bootstrap
	@echo "+ $@"
	@go mod tidy

bootstrap:
ifndef HAS_LINT
	GO111MODULE=off go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.20.0
endif

fmt:
	@echo "+ $@"
	@gofmt -l -s -w .

test:
	@echo "+ $@"
	@go test -count=1 -v -race -cover ./...

lint: bootstrap
	@echo "+ $@"
	@golangci-lint run ./...

.PHONY: all \
		fmt \
		vendor \
		bootstrap \
		lint
